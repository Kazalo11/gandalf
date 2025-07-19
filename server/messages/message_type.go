package messages

import (
	"encoding/json"
	"fmt"
)

type MessageType int

const (
	PlayerMessageType MessageType = iota
	GameMessageType
)

var (
	messageTypeToString = map[MessageType]string{
		PlayerMessageType: "PlayerMessage",
		GameMessageType:   "GameMessage",
	}
	stringToMessageType = map[string]MessageType{
		"PlayerMessage": PlayerMessageType,
		"GameMessage":   GameMessageType,
	}
)

func (m MessageType) MarshalJSON() ([]byte, error) {
	str, ok := messageTypeToString[m]
	if !ok {
		return nil, fmt.Errorf("invalid MessageType: %d", m)
	}
	return json.Marshal(str)
}

func (m *MessageType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	msgType, ok := stringToMessageType[str]
	if !ok {
		return fmt.Errorf("unknown message type: %s", str)
	}
	*m = msgType
	return nil
}
