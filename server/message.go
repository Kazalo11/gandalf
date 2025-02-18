package server

import (
	"encoding/json"
	"fmt"
)

type Action int

const (
	DrawCard Action = iota
	PlayCard
	Look
	ShowCard
)

func (a Action) String() string {
	return [...]string{"DrawCard", "PlayCard", "Look", "ShowCard"}[a]
}

func (a *Action) UnmarshalJSON(data []byte) error {
	var actionStr string
	if err := json.Unmarshal(data, &actionStr); err != nil {
		return err
	}

	switch actionStr {
	case "DrawCard":
		*a = DrawCard
	case "PlayCard":
		*a = PlayCard
	case "Look":
		*a = Look
	case "ShowCard":
		*a = ShowCard
	default:
		return fmt.Errorf("unknown action: %s", actionStr)
	}

	return nil
}

type Message struct {
	Action   Action `json:"action"`
	PlayerId string `json:"playerId"`
	Data     any    `json:"data"`
}

func parseMessage(message []byte) (Message, error) {
	var m Message
	err := json.Unmarshal(message, &m)
	if err != nil {
		return Message{}, err
	}
	return m, nil

}
