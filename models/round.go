package models

type Round struct {
	Turn      int
	IsGandalf bool
}

func (r *Round) NextTurn() {
	(*r).Turn -= 1
}
