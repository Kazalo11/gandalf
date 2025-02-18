package messages

import (
	"github.com/google/uuid"
)

type Message struct {
	Type     MessageType `json:"type"`
	PlayerID uuid.UUID   `json:"id"`
	Action   string      `json:"action"`
}
