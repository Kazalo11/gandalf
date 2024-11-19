package internals

import "github.com/Kazalo11/gandalf/models"

type Game struct {
	Deck    models.Deck
	Discard models.Discard
	Players []models.Player
	Rounds  []models.Round
}

func InitGame(numOfPlayers int) Game {
	g := Game{}
	g.Deck = models.InitDeck()
	g.Discard = models.Discard{}
	for i := 0; i < numOfPlayers; i++ {
		player := models.InitPlayer("Kazal")
		player.DrawCards(&g.Deck, 4)
		g.Players = append(g.Players, player)
	}
	return g

}
