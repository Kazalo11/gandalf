package models

type Deck []Card

func InitDeck() Deck {
	deck := make([]Card, 52)
	for i := 0; i <= len(deck)-1; i++ {
		deck[i].Rank = (i % 13) + 2
		deck[i].Suit = (i / 13)
	}

	return deck
}

func (d *Deck) Draw() Card {
	card := (*d)[0]
	(*d) = (*d)[1:]
	return card

}
