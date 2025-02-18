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

func broadcastReceivedMessages() {
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

func sendMessage(conn *websocket.Conn) {
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			break
		}
		fmt.Printf("Received %s \n", message)

		m, err := parseMessage(message)
		if err != nil {
			fmt.Printf("Error parsing message: %v\n", err)
			continue
		}
		fmt.Printf("Message parsed: %v\n", m)

		broadcast <- message
	}
}
