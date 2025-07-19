package server

import (
	"fmt"
	"github.com/Kazalo11/gandalf/models"
	"github.com/google/uuid"
)

type Hub struct {
	clients map[*Client]bool

	broadcast chan []byte

	register chan *Client

	unregister chan *Client

	game      *models.Game
	playerMap map[uuid.UUID]*Client
}

var (
	hubMap = make(map[uuid.UUID]*Hub)
)

func newHub(g *models.Game) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		game:       g,
		playerMap:  make(map[uuid.UUID]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.playerMap[client.player.Id] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				h.playerMap[client.player.Id] = nil
				close(client.send)
			}
		case message := <-h.broadcast:
			fmt.Printf("Receieved message at hub: %s\n", message)
			processMessage(message, h)
		}
	}
}
