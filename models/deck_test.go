package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDeck(t *testing.T) {
	deck := InitDeck()
	if len(deck) != 52 {
		t.Errorf("Expected deck size of 52, got %d", len(deck))
	}

	cardSet := make(map[string]bool)
	for _, card := range deck {
		key := fmt.Sprintf("%d-%d", card.Rank, card.Suit)
		if cardSet[key] {
			t.Errorf("Duplicate card found: Rank %d, Suit %d", card.Rank, card.Suit)
		}
		cardSet[key] = true
	}

	if len(cardSet) != 52 {
		t.Errorf("Expected 52 unique cards, got %d", len(cardSet))
	}
}

func TestDrawFromDeck(t *testing.T) {
	deck := InitDeck()
	initialSize := len(deck)

	card, err := deck.DrawFromDeck()
	if err != nil {
		t.Errorf("Error thrown %v", err)
	}
	if len(deck) != initialSize-1 {
		t.Errorf("Expected deck size %d after draw, got %d", initialSize-1, len(deck))
	}

	if card.Rank < ACE || card.Rank > KING {
		t.Errorf("Drawn card has invalid rank: %d", card.Rank)
	}

	if card.Suit < SPADE || card.Suit > HEART {
		t.Errorf("Drawn card has invalid suit: %d", card.Suit)
	}
}

func TestIsEmpty(t *testing.T) {
	deck := InitDeck()

	for i := 0; i < 52; i++ {
		_, err := deck.DrawFromDeck()
		if err != nil {
			t.Errorf("Threw an early error when trying to draw a card")
		}
	}

	if !deck.IsEmpty() {
		t.Errorf("Expected deck to be empty after drawing all cards")
	}

	_, err := deck.DrawFromDeck()
	assert.Error(t, err)

}
