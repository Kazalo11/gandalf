package server

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

type SubscribeMessage struct {
	Action string  `json:"action"`
	ID     *string `json:"id,omitempty"`
}

func handleMessage(msg []byte) {
	var parsedMsg SubscribeMessage
	if err := json.Unmarshal(msg, &parsedMsg); err != nil {
		log.Printf("Failed to parse message: %v", err)
		return
	}

	action, id := parsedMsg.Action, parsedMsg.ID

	if id != nil {
		ID := uuid.MustParse(*id)
	}

}
