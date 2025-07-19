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

type GetGameState struct {
	GameBaseMessage
	GameId uuid.UUID `json:"gameId"`
}

func parseGameMessage(message []byte) (Message, error) {
	var m GameBaseMessage
	err := json.Unmarshal(message, &m)
	if err != nil {
		return &m, err
	}
	switch m.MessageSubType {
	case GetGameStateMessage:
		var getState GetGameState
		if err := json.Unmarshal(message, &getState); err != nil {
			return &m, err
		}
		return &getState, nil
	case JoinGameMessage:
		var join JoinGame
		if err := json.Unmarshal(message, &join); err != nil {
			return &m, err
		}
		return &join, nil
	}
	return &m, nil
}
