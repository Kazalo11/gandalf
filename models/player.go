package models

import "fmt"

type Player struct {
	Name string
	Hand Deck
}

func (p *Player) DrawCards(d *Deck, numberOfCards int) {
	for i := 0; i < numberOfCards; i++ {
		p.Hand = append(p.Hand, (*d).Draw())

	}
	fmt.Printf("Cards drawn: %v \n", p.Hand)
}
