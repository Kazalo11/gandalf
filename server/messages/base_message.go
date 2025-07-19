package messages

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type MessageType int

const (
	PlayerMessageType MessageType = iota
	GameMessageType
)

var (
	messageTypeToString = map[MessageType]string{
		PlayerMessageType: "PlayerMessage",
		GameMessageType:   "GameBaseMessage",
	}
	stringToMessageType = map[string]MessageType{
		"PlayerMessage":   PlayerMessageType,
		"GameBaseMessage": GameMessageType,
	}
)

func (m *MessageType) MarshalJSON() ([]byte, error) {
	str, ok := messageTypeToString[*m]
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

type Message interface {
	GetType() MessageType
}

type BaseMessage struct {
	Id          uuid.UUID   `json:"id"`
	MessageType MessageType `json:"type"`
}

func (m *BaseMessage) GetType() MessageType {
	return m.MessageType
}

func ParseMessage(message []byte) (Message, error) {
	var base BaseMessage
	if err := json.Unmarshal(message, &base); err != nil {
		return nil, err
	}

	switch base.MessageType {
	case PlayerMessageType:
		msg, err := parsePlayerMessage(message)
		if err != nil {
			return nil, err
		}
		return &msg, nil
	case GameMessageType:
		msg, err := parseGameMessage(message)
		if err != nil {
			return nil, err
		}
		return &msg, nil
	default:
		return nil, errors.New("unknown message type")
	}
}
