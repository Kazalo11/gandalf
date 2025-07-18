package internals

import (
	"github.com/Kazalo11/gandalf/models"
	"github.com/google/uuid"
)

type Game struct {
	Deck    models.Deck
	Discard models.Discard
	Players map[uuid.UUID]*models.Player
	Rounds  []models.Round
	Id      uuid.UUID
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
