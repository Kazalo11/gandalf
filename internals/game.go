package internals

import "github.com/Kazalo11/gandalf/models"

type Game struct {
	Deck    models.Deck
	Discard models.Discard
	Players []models.Player
	Rounds  []models.Round
}

func InitGame() *Game {
	g := Game{}
	g.Deck = models.InitDeck()
	return &g
}

func (g *Game) AddPlayer(p models.Player) {
	g.Players = append(g.Players, p)
}
