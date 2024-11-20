package models

import (
	"fmt"

	"log"

	"github.com/google/uuid"
)

type Player struct {
	Name string
	Hand Deck
	Id   uuid.UUID
}

type UUIDGenerator func() uuid.UUID

var DefaultUUIDGenerator UUIDGenerator = uuid.New

func InitPlayer(name string) Player {
	return Player{
		Name: name,
		Hand: Deck{},
		Id:   DefaultUUIDGenerator(),
	}
}

func (p *Player) DrawCards(d *Deck, numberOfCards int) {
	for i := 0; i < numberOfCards; i++ {
		card, err := (*d).DrawFromDeck()
		if err != nil {
			log.Printf("Couldn't withdraw a card due to: %v", err)
		} else {

			p.Hand = append(p.Hand, card)
		}

	}
	fmt.Printf("Cards drawn: %v \n", p.Hand)
}

func (p *Player) Draw(d *Deck) Card {
	card, err := (*d).DrawFromDeck()
	if err != nil {
		log.Printf("Couldn't withdraw a card due to: %v", err)
	}
	return card
}

func (p *Player) CalculateScore() int {
	score := 0
	for _, card := range p.Hand {
		score += card.Value()

	}
	return score
}

func (p *Player) Look(cardNum int) (string, string) {
	return p.Hand[cardNum].Show()
}

func (p *Player) PlayCard(cardNum int, d *Discard) {
	card := p.Hand[cardNum]
	p.Hand = append(p.Hand[:cardNum], p.Hand[cardNum+1:]...)
	(*d).AddCard(card)
}

func (p *Player) SwapCards(p2 *Player, cardNum1, cardNum2 int) {
	p.Hand[cardNum1], p2.Hand[cardNum2] = p2.Hand[cardNum1], p.Hand[cardNum1]
}

func (p *Player) CallGandalf(g *Game) {
	finalRound := Round{
		Turns:       len(g.Players) - 1,
		IsGandalf:   true,
		CurrentTurn: g.Round.CurrentTurn,
	}
	g.Round = finalRound
}
