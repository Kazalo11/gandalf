package models

import (
	"github.com/google/uuid"
)

type Game struct {
	Deck    Deck                  `json:"deck"`
	Discard Discard               `json:"discard"`
	Players map[uuid.UUID]*Player `json:"players"`
	Rounds  []Round               `json:"rounds"`
	Id      uuid.UUID             `json:"id"`
}

func InitGame(gameId uuid.UUID) *Game {
	g := Game{}
	g.Id = gameId
	g.Deck = InitDeck()
	g.Discard = make([]Card, 0)
	g.Rounds = make([]Round, 0)
	g.Rounds = append(g.Rounds, Round{Turn: 1, IsGandalf: true})
	g.Players = make(map[uuid.UUID]*Player)
	return &g
}

func (g *Game) AddPlayer(p Player) {
	g.Players[p.Id] = &p
}
