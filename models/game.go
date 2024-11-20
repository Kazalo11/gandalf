package models

import (
	"errors"

	"github.com/google/uuid"
)

type Game struct {
	Deck    Deck
	Discard Discard
	Players []Player
	Rounds  []Round
}

func InitGame(numOfPlayers int) Game {
	g := Game{}
	g.Deck = InitDeck()
	g.Discard = Discard{}
	for i := 0; i < numOfPlayers; i++ {
		player := InitPlayer("Kazal")
		player.DrawCards(&g.Deck, 4)
		g.Players = append(g.Players, player)
	}
	return g

}

func (g *Game) GetPlayer(id uuid.UUID) (Player, error) {
	for _, player := range g.Players {
		if id == player.Id {
			return player, nil
		}
	}
	return Player{}, errors.New("couldn't find that player in the game")

}
