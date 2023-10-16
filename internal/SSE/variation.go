package sse

import (
	"net/http"
)

type Broker struct {
	Notifier       chan []byte
	NewClients     chan chan []byte
	ClosingClients chan chan []byte
	Clients        map[chan []byte]bool
}

func NewBroker() *Broker {
	return &Broker{
		Notifier:       make(chan []byte),
		NewClients:     make(chan chan []byte),
		ClosingClients: make(chan chan []byte),
		Clients:        make(map[chan []byte]bool),
	}
}

func (b *Broker) Listen() {
	for {
		select {
		case s := <-b.NewClients:
			b.Clients[s] = true
		case s := <-b.ClosingClients:
			delete(b.Clients, s)
			close(s)
		case event := <-b.Notifier:
			for clientMessageChan := range b.Clients {
				select {
				case clientMessageChan <- event:
				default:
					close(clientMessageChan)
					delete(b.Clients, clientMessageChan)
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

	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		http.Error(w, "User ID not found!", http.StatusBadRequest)
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	messageChan := make(chan []byte, 1)
	broker.NewClients <- messageChan

	defer func() {
		broker.ClosingClients <- messageChan
	}()

	notify := w.(http.CloseNotifier).CloseNotify()

	go func() {
		<-notify
		broker.ClosingClients <- messageChan
	}()

	for msg := range messageChan {
		w.Write(msg)
		flusher.Flush()
	}
}

func Teste(w http.ResponseWriter, r *http.Request) {
	broker.Notifier <- []byte("Hello, world!")

	w.Write([]byte("OK"))
}
