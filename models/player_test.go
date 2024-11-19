package models

import (
	"testing"
)

func TestPlayer_DrawCards(t *testing.T) {
	deck := InitDeck()
	player := Player{Name: "Alice"}

	player.DrawCards(&deck, 5)
	if len(player.Hand) != 5 {
		t.Errorf("Expected player hand size to be 5, got %d", len(player.Hand))
	}

	if len(deck) != 47 {
		t.Errorf("Expected deck size to be 47 after drawing 5 cards, got %d", len(deck))
	}
}

func TestPlayer_Draw(t *testing.T) {
	deck := InitDeck()
	player := Player{Name: "Bob"}

	card := player.Draw(&deck)
	if len(deck) != 51 {
		t.Errorf("Expected deck size to be 51 after drawing a card, got %d", len(deck))
	}

	if card.Rank < ACE || card.Rank > KING {
		t.Errorf("Drawn card has invalid rank: %d", card.Rank)
	}

	if card.Suit < SPADE || card.Suit > HEART {
		t.Errorf("Drawn card has invalid suit: %d", card.Suit)
	}
}

func TestPlayer_CalculateScore(t *testing.T) {
	player := Player{Name: "Alice", Hand: Deck{
		{Rank: ACE, Suit: SPADE},
		{Rank: TEN, Suit: HEART},
		{Rank: KING, Suit: CLUB},
	}}

	expectedScore := 21
	score := player.CalculateScore()
	if score != expectedScore {
		t.Errorf("Expected score %d, got %d", expectedScore, score)
	}
}

func TestPlayer_Look(t *testing.T) {
	player := Player{Name: "Charlie", Hand: Deck{
		{Rank: QUEEN, Suit: HEART},
	}}

	rank, suit := player.Look(0)
	if rank != "Q" || suit != "♥" {
		t.Errorf("Expected card Q♥, got %s%s", rank, suit)
	}
}

func TestPlayer_PlayCard(t *testing.T) {
	player := Player{Name: "Dave", Hand: Deck{
		{Rank: THREE, Suit: DIAMOND},
		{Rank: FIVE, Suit: SPADE},
	}}
	discard := Discard{}

	player.PlayCard(0, &discard)
	if len(player.Hand) != 1 {
		t.Errorf("Expected hand size 1, got %d", len(player.Hand))
	}

	if len(discard) != 1 {
		t.Errorf("Expected discard pile size 1, got %d", len(discard))
	}

	if discard[0].Rank != THREE || discard[0].Suit != DIAMOND {
		t.Errorf("Expected discarded card to be 3♦, got %d %d", discard[0].Rank, discard[0].Suit)
	}
}

func TestPlayer_SwapCards(t *testing.T) {
	player1 := Player{Name: "Alice", Hand: Deck{
		{Rank: ACE, Suit: SPADE},
	}}
	player2 := Player{Name: "Bob", Hand: Deck{
		{Rank: KING, Suit: HEART},
	}}

	player1.SwapCards(&player2, 0, 0)
	if player1.Hand[0].Rank != KING || player1.Hand[0].Suit != HEART {
		t.Errorf("Expected Alice's card to be K♥, got %d %d", player1.Hand[0].Rank, player1.Hand[0].Suit)
	}

	if player2.Hand[0].Rank != ACE || player2.Hand[0].Suit != SPADE {
		t.Errorf("Expected Bob's card to be A♠, got %d %d", player2.Hand[0].Rank, player2.Hand[0].Suit)
	}
}
