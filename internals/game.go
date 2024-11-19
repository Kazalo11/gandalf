package internals

import "github.com/Kazalo11/gandalf/models"

type Game struct {
	Deck    models.Deck
	Discard models.Discard
	Players []models.Player
	Rounds  []models.Round
}

func (g *Game) InitGame(numOfPlayers int) {
	g.Deck = models.InitDeck()
	for i := 0; i < numOfPlayers; i++ {

	}

}
