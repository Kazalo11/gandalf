package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Kazalo11/gandalf/models"
	"github.com/Kazalo11/gandalf/server/messages"
	"log"
	"net/http"
	"time"

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
		err := c.conn.Close()
		if err != nil {
			return
		}
	}()
	c.conn.SetReadLimit(maxMessageSize)
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		return
	}
	c.conn.SetPongHandler(func(string) error {
		err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			return err
		}
		return nil
	})
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
		_, err = messages.ParseMessage(message)
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
		err := c.conn.Close()
		if err != nil {
			return
		}
	}()
	for {
		select {
		case message, ok := <-c.send:
			err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				return
			}
			if !ok {
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					return
				}
				return
			}

			err = c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}

		case <-ticker.C:
			err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				return
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func connectToHub(hub *Hub, p *models.Player, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connecting to hub")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	joinGameMessage := messages.JoinGame{
		GameBaseMessage: messages.GameBaseMessage{
			BaseMessage: messages.BaseMessage{
				MessageType:    messages.GameMessageType,
				Id:             hub.game.Id,
				MessageSubType: messages.JoinGameMessage,
			},
		},
		PlayerId: p.Id,
		GameId:   hub.game.Id,
	}

	msg, err := json.Marshal(joinGameMessage)
	if err == nil {
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), player: p}

	go func() {
		client.hub.register <- client
	}()
	fmt.Println("Registering client in hub")

	go client.receiveMessages()
	go client.sendMessages()

}
