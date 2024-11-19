package models

type Discard []Card

func (d *Discard) AddCard(c Card) {
	(*d) = append((*d), c)
}

func (d *Discard) IsEmpty() bool {
	return len(*d) == 0
}

func (d *Discard) DrawFromDiscard() Card {
	card := (*d)[len(*d)-1]
	(*d) = (*d)[:len(*d)-1]
	return card
}
