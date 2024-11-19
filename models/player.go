package models

import "fmt"

type Player struct {
	Name string
	Hand Deck
}

func (p *Player) DrawCards(d *Deck, numberOfCards int) {
	for i := 0; i < numberOfCards; i++ {
		p.Hand = append(p.Hand, (*d).DrawFromDeck())

	}
	fmt.Printf("Cards drawn: %v \n", p.Hand)
}

func (p *Player) Draw(d *Deck) Card {
	return (*d).DrawFromDeck()
}

func (p *Player) CalculateScore() int {
	score := 0
	for _, card := range p.Hand {
		score += card.Value()

	}
	return score
}
