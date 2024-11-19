package models

import "testing"

func TestNextTurn(t *testing.T) {
	tests := []struct {
		initialTurn  int
		expectedTurn int
	}{
		{initialTurn: 5, expectedTurn: 4},
		{initialTurn: 1, expectedTurn: 0},
		{initialTurn: 0, expectedTurn: 0},
	}

	for _, tt := range tests {
		t.Run("Testing NextTurn", func(t *testing.T) {
			round := Round{Turn: tt.initialTurn}
			round.NextTurn()

			if round.Turn != tt.expectedTurn {
				t.Errorf("Expected Turn %d, got %d", tt.expectedTurn, round.Turn)
			}
		})
	}
}

func TestSkipTurn(t *testing.T) {
	tests := []struct {
		initialTurn  int
		expectedTurn int
	}{
		{initialTurn: 5, expectedTurn: 3},
		{initialTurn: 2, expectedTurn: 0},
		{initialTurn: 1, expectedTurn: 0},
	}

	for _, tt := range tests {
		t.Run("Testing SkipTurn", func(t *testing.T) {
			round := Round{Turn: tt.initialTurn}
			round.SkipTurn()

			if round.Turn != tt.expectedTurn {
				t.Errorf("Expected Turn %d, got %d", tt.expectedTurn, round.Turn)
			}
		})
	}
}
