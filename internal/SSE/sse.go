package sse

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Broker struct {
	Notifier         chan []byte
	NewClients       chan ClientInfo
	ClosingClients   chan ClientInfo
	Clients          map[string]map[chan []byte]bool // Map by userID
	TargetedNotifier chan Notification
}

type ClientInfo struct {
	UserID string
	Chan   chan []byte
}

type Notification struct {
	UserID  string
	Message []byte
}

func NewBroker() *Broker {
	return &Broker{
		Notifier:         make(chan []byte),
		NewClients:       make(chan ClientInfo),
		ClosingClients:   make(chan ClientInfo),
		Clients:          make(map[string]map[chan []byte]bool),
		TargetedNotifier: make(chan Notification),
	}
}

func (b *Broker) Listen() {
	for {
		select {
		case client := <-b.NewClients:
			if _, ok := b.Clients[client.UserID]; !ok {
				b.Clients[client.UserID] = make(map[chan []byte]bool)
			}
			b.Clients[client.UserID][client.Chan] = true
		case client := <-b.ClosingClients:
			if chans, ok := b.Clients[client.UserID]; ok {
				delete(chans, client.Chan)
				if len(chans) == 0 {
					delete(b.Clients, client.UserID)
				}
			}
		case event := <-b.Notifier:
			for _, clientChans := range b.Clients {
				for clientMessageChan := range clientChans {
					clientMessageChan <- event
				}
			}
		case notification := <-b.TargetedNotifier:
			if clientChans, ok := b.Clients[notification.UserID]; ok {
				for clientMessageChan := range clientChans {
					clientMessageChan <- notification.Message
				}
			}
		}
	}
}

var broker = NewBroker()

func init() {
	go broker.Listen()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	userID := mux.Vars(r)["user_id"]
	if userID == "" {
		http.Error(w, "User ID not provided!", http.StatusBadRequest)
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	closeNotifier, ok := w.(http.CloseNotifier)
	if !ok {
		http.Error(w, "Close notification unsupported!", http.StatusInternalServerError)
		return
	}

	messageChan := make(chan []byte, 1)
	clientInfo := ClientInfo{UserID: userID, Chan: messageChan}
	broker.NewClients <- clientInfo

	defer func() {
		broker.ClosingClients <- clientInfo
	}()

	notify := closeNotifier.CloseNotify()
	go func() {
		<-notify
		broker.ClosingClients <- clientInfo
	}()

	for msg := range messageChan {
		w.Write(msg)
		flusher.Flush()
	}
}
