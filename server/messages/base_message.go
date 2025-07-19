package messages

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Message interface {
	GetType() MessageType
	GetSubType() MessageSubType
}

type BaseMessage struct {
	Id             uuid.UUID      `json:"id"`
	MessageType    MessageType    `json:"type"`
	MessageSubType MessageSubType `json:"subtype"`
}

func (m *BaseMessage) GetType() MessageType {
	return m.MessageType
}

func (m *BaseMessage) GetSubType() MessageSubType {
	return m.MessageSubType
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
		return msg, nil
	default:
		return nil, errors.New("unknown message type")
	}
}
