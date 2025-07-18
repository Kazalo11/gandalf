package server

import "github.com/google/uuid"

type GameResponse struct {
	GameID   uuid.UUID `json:"gameId"`
	PlayerID uuid.UUID `json:"playerId"`
}
