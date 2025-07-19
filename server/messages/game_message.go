package messages

import (
	"encoding/json"
	"github.com/google/uuid"
)

type GameBaseMessage struct {
	BaseMessage
}

type JoinGame struct {
	GameBaseMessage
	PlayerId uuid.UUID `json:"playerId"`
	GameId   uuid.UUID `json:"gameId"`
}

func parseGameMessage(message []byte) (Message, error) {
	var m GameBaseMessage
	err := json.Unmarshal(message, &m)
	if err != nil {
		return &m, err
	}
	if m.MessageSubType == JoinGameMessage {
		var join JoinGame
		if err := json.Unmarshal(message, &join); err != nil {
			return &m, err
		}
		return &join, nil
	}
	return &m, nil
}
