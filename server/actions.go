package server

import (
	"fmt"

	"github.com/Kazalo11/gandalf/internals"
	"github.com/Kazalo11/gandalf/models"
	"github.com/google/uuid"
)

func processMessage(m Message, g *internals.Game) models.Card {
	players := g.Players
	playerId, err := uuid.Parse(m.PlayerId)
	if err != nil {
		fmt.Printf("%s is not a valid player Id", m.PlayerId)
		return models.Card{}
	}
	player := players[playerId]

	switch m.Action {
	case DrawCard:
		fmt.Println("Drawing card")
		card := player.Draw(&g.Deck)
		fmt.Printf("Card drawn is %v", card)
		return card
	case PlayCard:
		fmt.Println("Playing card")
	case Look:
		fmt.Println("Looking at card")
	case ShowCard:
		fmt.Println("Showing card")
	default:
		fmt.Printf("%s is not a valid action", m.Action.String())
		return models.Card{}
	}
	return models.Card{}
}
