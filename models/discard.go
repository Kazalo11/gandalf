package models

type Discard []Card

func (d *Discard) AddCard(c Card) {
	(*d) = append((*d), c)
}

func (d *Discard) IsEmpty() bool {
	return len(*d) == 0
}

func (d *Discard) DrawFromDiscard() Card {
	card := (*d)[0]
	(*d) = (*d)[1:]
	return card
}
