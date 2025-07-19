package messages

import (
	"encoding/json"
	"fmt"
)

type MessageSubType int

const (
	JoinGameMessage MessageSubType = iota
	GetGameStateMessage
	GameStateMessage
)

var (
	messageSubTypeToString = map[MessageSubType]string{
		JoinGameMessage:     "JoinGame",
		GetGameStateMessage: "GetGame",
		GameStateMessage:    "GameState",
	}

	stringToMessageSubType = map[string]MessageSubType{
		"JoinGame":  JoinGameMessage,
		"GetGame":   GetGameStateMessage,
		"GameState": GameStateMessage,
	}
)

func (m MessageSubType) MarshalJSON() ([]byte, error) {
	str, ok := messageSubTypeToString[m]
	if !ok {
		return nil, fmt.Errorf("invalid MessageSubType: %d", m)
	}
	return json.Marshal(str)
}

func (m *MessageSubType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	msgSubType, ok := stringToMessageSubType[str]
	if !ok {
		return fmt.Errorf("unknown message subtype: %s", str)
	}
	*m = msgSubType
	return nil
}
