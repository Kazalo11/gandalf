package messages

import (
	"encoding/json"
)

type GameMessage struct {
	BaseMessage
	Data any `json:"data"`
}

func parseGameMessage(message []byte) (GameMessage, error) {
	var m GameMessage
	err := json.Unmarshal(message, &m)
	if err != nil {
		return GameMessage{}, err
	}
	return m, nil
}
