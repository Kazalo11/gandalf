package messages

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

type MessageType int

const (
	PlayerMessageType MessageType = iota
	GameMessageType
)

type Message interface {
	GetType() MessageType
}

type BaseMessage struct {
	Id          uuid.UUID `json:"id"`
	MessageType string    `json:"type"`
}

func (m *BaseMessage) GetType() MessageType {
	switch m.MessageType {
	case "PlayerMessage":
		return PlayerMessageType
	case "GameMessage":
		return GameMessageType
	default:
		return -1
	}
}

func ParseMessage(message []byte) (Message, error) {
	var base BaseMessage
	if err := json.Unmarshal(message, &base); err != nil {
		return nil, err
	}

	switch base.MessageType {
	case "PlayerMessage":
		playerMessage, err := parsePlayerMessage(message)
		if err != nil {
			return nil, err
		}
		return &playerMessage, nil
	case "GameMessage":
		gameMessage, err := parseGameMessage(message)
		if err != nil {
			return nil, err
		}
		return &gameMessage, nil
	default:
		return nil, errors.New("unknown message type")
	}
}
