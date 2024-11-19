package models

import (
	"fmt"
	"testing"
)

func TestCardValue(t *testing.T) {
	tests := []struct {
		card     Card
		expected int
	}{
		{Card{Suit: SPADE, Rank: ACE}, 1},
		{Card{Suit: CLUB, Rank: TWO}, 2},
		{Card{Suit: DIAMOND, Rank: TEN}, 10},
		{Card{Suit: HEART, Rank: JACK}, 10},
		{Card{Suit: SPADE, Rank: QUEEN}, 10},
		{Card{Suit: CLUB, Rank: KING}, 10},
		{Card{Suit: DIAMOND, Rank: 0}, 0},
	}

	for _, test := range tests {
		rank, suit := test.card.Show()
		t.Run(fmt.Sprintf("Testing Value of %s %s \n", rank, suit), func(t *testing.T) {
			result := test.card.Value()
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestCardShow(t *testing.T) {
	tests := []struct {
		card         Card
		expectedRank string
		expectedSuit string
	}{
		{Card{Suit: SPADE, Rank: ACE}, "A", "♠"},
		{Card{Suit: HEART, Rank: KING}, "K", "♥"},
		{Card{Suit: CLUB, Rank: QUEEN}, "Q", "♣"},
		{Card{Suit: DIAMOND, Rank: JACK}, "J", "♦"},
		{Card{Suit: CLUB, Rank: TWO}, "2", "♣"},
		{Card{Suit: DIAMOND, Rank: NINE}, "9", "♦"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Testing Show of %v", test.card), func(t *testing.T) {
			rank, suit := test.card.Show()
			if rank != test.expectedRank || suit != test.expectedSuit {
				t.Errorf("Expected rank: %s, suit: %s, got rank: %s, suit: %s", test.expectedRank, test.expectedSuit, rank, suit)
			}
		})
	}
}
