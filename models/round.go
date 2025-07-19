package models

type Round struct {
	Turn      int  `json:"turn"`
	IsGandalf bool `json:"isGandalf"`
}

func (r *Round) NextTurn() {
	if r.Turn == 0 {
		return
	}
	(*r).Turn -= 1

}

func (r *Round) SkipTurn() {
	if r.Turn < 2 {
		(*r).Turn = 0
	} else {
		(*r).Turn -= 2
	}
}
