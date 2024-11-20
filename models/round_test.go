package models

import "testing"

func TestNextTurn(t *testing.T) {
	tests := []struct {
		initialTurn  int
		expectedTurn int
	}{
		{initialTurn: 2, expectedTurn: 3},
		{initialTurn: 1, expectedTurn: 2},
		{initialTurn: 3, expectedTurn: 0},
	}

	for _, tt := range tests {
		t.Run("Testing NextTurn", func(t *testing.T) {
			round := Round{CurrentTurn: tt.initialTurn, Turns: 4}
			round.NextTurn()

			if round.CurrentTurn != tt.expectedTurn {
				t.Errorf("Expected Turn %d, got %d", tt.expectedTurn, round.CurrentTurn)
			}
		})
	}
}

func TestSkipTurn(t *testing.T) {
	tests := []struct {
		initialTurn  int
		expectedTurn int
	}{
		{initialTurn: 2, expectedTurn: 4},
		{initialTurn: 1, expectedTurn: 3},
		{initialTurn: 3, expectedTurn: 0},
	}

	for _, tt := range tests {
		t.Run("Testing SkipTurn", func(t *testing.T) {
			round := Round{CurrentTurn: tt.initialTurn, Turns: 5}
			round.SkipTurn()

			if round.CurrentTurn != tt.expectedTurn {
				t.Errorf("Expected Turn %d, got %d", tt.expectedTurn, round.CurrentTurn)
			}
		})
	}
}
