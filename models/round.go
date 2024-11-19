package models

type Round struct {
	Turn      int
	IsGandalf bool
}

func (r *Round) NextTurn() {
	if r.Turn == 0 {
		return
	}
	(*r).Turn -= 1

}
