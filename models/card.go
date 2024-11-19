package models

const (
	SPADE = iota
	CLUB
	DIAMOND
	HEART
)

const (
	ACE = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

type Card struct {
	Suit int
	Rank int
}

func (c Card) Value() int {
	switch c.Rank {
	case ACE:
		return 1
	case TWO:
		return 2
	case THREE:
		return 3
	case FOUR:
		return 4
	case FIVE:
		return 5
	case SIX:
		return 6
	case SEVEN:
		return 7
	case EIGHT:
		return 8
	case NINE:
		return 9
	case TEN:
		return 10
	case JACK:
		return 10
	case QUEEN:
		return 10
	case KING:
		return 10
	default:
		return 0
	}
}

func (c Card) toSuit() string {
	switch c.Suit {
	case SPADE:
		return "♠"
	case HEART:
		return "♥"
	case CLUB:
		return "♣"
	case DIAMOND:
		return "♦"
	default:
		return ""
	}
}
