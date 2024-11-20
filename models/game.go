package models

import (
	"errors"

	"github.com/google/uuid"
)

type Game struct {
	Deck    Deck
	Discard Discard
	Players []Player
	Round   Round
}

func InitGame(numOfPlayers int) Game {
	g := Game{}
	g.Deck = InitDeck()
	g.Discard = Discard{}
	g.Round = Round{
		Turns:     numOfPlayers,
		IsGandalf: false,
	}
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

func (g *Game) EndGame() (bool, Player) {
	maxScore := 0
	var winningPlayer Player
	if !g.Round.IsGandalf {
		return false, winningPlayer
	}

	for _, player := range g.Players {
		score := player.CalculateScore()
		if score > maxScore {
			maxScore = score
			winningPlayer = player
		}

	}
	return true, winningPlayer

}
