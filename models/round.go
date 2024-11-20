package models

type Round struct {
	Turns       int
	IsGandalf   bool
	CurrentTurn int
}

func (r *Round) NextTurn() {

	if r.CurrentTurn == r.Turns-1 {
		new_round := Round{
			Turns:       r.Turns,
			IsGandalf:   false,
			CurrentTurn: (r.CurrentTurn + 1) % r.Turns,
		}
		(*r) = new_round
	} else {

		(*r).CurrentTurn += 1
	}

}

func (r *Round) SkipTurn() {
	if (r.Turns - r.CurrentTurn) <= 2 {
		new_round := Round{
			Turns:       r.Turns,
			IsGandalf:   false,
			CurrentTurn: (r.CurrentTurn + 2) % r.Turns,
		}
		(*r) = new_round
	} else {

		(*r).CurrentTurn += 2
	}

}
