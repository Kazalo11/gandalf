package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Kazalo11/gandalf/models"
	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
	mutex   sync.Mutex
)

type Client struct {
	hub *Hub

	conn *websocket.Conn

	send chan []byte

	player *models.Player
}

func (c *Client) sendMessages() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		fmt.Printf("Received message: %s\n", message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		_, err = parseMessage(message)
		if err != nil {
			fmt.Printf("Can't parse message due to %v, not sending\n", err)
			continue

		}

		c.hub.broadcast <- message

	}
}

func (c *Client) receiveMessages() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.WriteMessage(websocket.TextMessage, message)

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func connectToHub(hub *Hub, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connecting to hub")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	game := hub.game
	mutex.Lock()
	player := createPlayer(game)

	if player == nil {
		fmt.Println("Failed to create player")
		conn.Close()
		return
	}
	fmt.Println(*player)
	game.AddPlayer(*player)
	mutex.Unlock()

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), player: player}
	go func() {
		client.hub.register <- client
	}()
	fmt.Println("Registering client in hub")

	go client.receiveMessages()
	go client.sendMessages()

}
