package internals

import (
	"github.com/Kazalo11/gandalf/models"
	"github.com/google/uuid"
)

type Game struct {
	Deck    models.Deck                  `json:"deck"`
	Discard models.Discard               `json:"discard"`
	Players map[uuid.UUID]*models.Player `json:"players"`
	Rounds  []models.Round               `json:"rounds"`
	Id      uuid.UUID                    `json:"id"`
}

func InitGame(gameId uuid.UUID) *Game {
	g := Game{}
	g.Id = gameId
	g.Deck = models.InitDeck()
	g.Players = make(map[uuid.UUID]*models.Player)
	return &g
}

func (g *Game) AddPlayer(p models.Player) {
	g.Players[p.Id] = &p
}
