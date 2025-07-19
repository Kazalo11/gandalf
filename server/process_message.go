package server

import (
	"encoding/json"
	"fmt"
	"github.com/Kazalo11/gandalf/server/messages"
)

func processMessage(m []byte, h *Hub) {
	parsedMessage, err := messages.ParseMessage(m)
	if err != nil {
		fmt.Printf("Unable to parse message due to %v\n", err)
		return
	}

	fmt.Printf("Parsed message received: %+v\n", parsedMessage)

	switch parsedMessage.GetSubType() {
	case messages.JoinGameMessage:
		fmt.Printf("Received JoinMessage, updating game state for rest of the players\n")
		handleUpdateGameState(h)
		return
	case messages.GetGameStateMessage:
		fmt.Printf("Received GetGameStateMessage, updating game state for rest of the players\n")
		handleUpdateGameState(h)
		return
	default:
		fmt.Printf("Unknown message type: %s\n", parsedMessage.GetSubType())
	}
}

func handleUpdateGameState(h *Hub) {
	fmt.Println("Updating game state for all players")
	message := messages.GameState{
		Game: *h.game,
		GameBaseMessage: messages.GameBaseMessage{
			BaseMessage: messages.BaseMessage{
				Id:             h.game.Id,
				MessageType:    messages.GameMessageType,
				MessageSubType: messages.GameStateMessage,
			},
		},
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("Failed to marshal game state message: %v\n", err)
		return
	}

	for client := range h.clients {
		select {
		case client.send <- messageBytes:
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}
}
