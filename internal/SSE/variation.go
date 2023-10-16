package sse

import "net/http"

type Broker struct {
	Notifier       chan []byte
	NewClients     chan chan []byte
	ClosingClients chan chan []byte
	Clients        map[chan []byte]bool
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
				clientMessageChan <- event
			}
		}
	}
}

func Handler(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	messageChan := make(chan []byte)

	broker := &Broker{
		Notifier:       messageChan,
		NewClients:     make(chan chan []byte),
		ClosingClients: make(chan chan []byte),
		Clients:        make(map[chan []byte]bool),
	}

	go broker.Listen()

	messageChan <- []byte("data: Connected\n\n")

	flusher.Flush()
}
