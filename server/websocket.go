package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	broadcast = make(chan []byte)
	mutex     = &sync.Mutex{}
	clients   = []*websocket.Conn{}
)

func handleMessages() {
	for {
		message := <-broadcast

		mutex.Lock()
		fmt.Printf("Broadcasting message: %s \n", message)

		var activeClients []*websocket.Conn
		for _, client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				client.Close()
				continue
			}
			activeClients = append(activeClients, client)
		}

		clients = activeClients
		mutex.Unlock()
	}
}
