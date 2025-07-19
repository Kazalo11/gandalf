package messages

import (
	"encoding/json"
	"github.com/google/uuid"
)

type GameMessageSubType int

const (
	JoinGameMessage GameMessageSubType = iota
	GetGameMessage
)

type GameBaseMessage struct {
	BaseMessage
	SubType GameMessageSubType `json:"subtype"`
}

type JoinGame struct {
	GameBaseMessage
	PlayerId uuid.UUID `json:"playerId"`
	GameId   uuid.UUID `json:"gameId"`
}

func parseGameMessage(message []byte) (GameBaseMessage, error) {
	var m GameBaseMessage
	err := json.Unmarshal(message, &m)
	if err != nil {
		return GameBaseMessage{}, err
	}
	return m, nil
}
