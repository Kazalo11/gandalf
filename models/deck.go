package models

import (
	"errors"
	"fmt"

	"golang.org/x/exp/rand"
)

type Deck []Card

func InitDeck() Deck {
	deck := make([]Card, 52)
	for i := 0; i <= len(deck)-1; i++ {
		deck[i].Rank = (i % 13) + 1
		deck[i].Suit = (i / 13)
	}

	rand.Shuffle(52, func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	fmt.Println(deck)
	return deck
}

func (d *Deck) DrawFromDeck() (Card, error) {
	if d.IsEmpty() {
		return Card{0, 0}, errors.New("Can't draw a card, deck is empty")
	}
	card := (*d)[len(*d)-1]
	(*d) = (*d)[:len(*d)-1]
	return card, nil

}

func (d *Deck) IsEmpty() bool {
	return len(*d) == 0
}
