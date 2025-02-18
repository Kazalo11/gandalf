package server

import (
	"fmt"

	"github.com/Kazalo11/gandalf/internals"
	"github.com/google/uuid"
)

type Hub struct {
	clients map[*Client]bool

	broadcast chan []byte

	register chan *Client

	unregister chan *Client

	game *internals.Game
}

var (
	hubMap = make(map[uuid.UUID]*Hub)
)

func newHub(g *internals.Game) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		game:       g,
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			fmt.Printf("Receieved message at hub: %s\n", message)

			parsedMessage, err := parseMessage(message)
			if err != nil {
				fmt.Printf("Unable to parse error due to %v", err)
				continue
			}
			fmt.Printf("Parsed message receieved %s", parsedMessage)
			processMessage(parsedMessage, h.game)
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
