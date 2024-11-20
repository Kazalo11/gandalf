package server

import (
	"encoding/json"
	"log"
)

type Message struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

func handleMessage(msg []byte) {
	var parsedMsg Message
	if err := json.Unmarshal(msg, &parsedMsg); err != nil {
		log.Printf("Failed to parse message: %v", err)
		return
	}
}
