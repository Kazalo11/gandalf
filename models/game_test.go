package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInitGame(t *testing.T) {
	numOfPlayers := 3
	game := InitGame(numOfPlayers)

	assert.Equal(t, numOfPlayers, len(game.Players), "Game should have the correct number of players")

	expectedDeckSize := 52 - (numOfPlayers * 4)
	assert.Equal(t, expectedDeckSize, len(game.Deck), "Deck should have the correct number of remaining cards")

	assert.Equal(t, 0, len(game.Discard), "Discard pile should be empty at initialization")

	assert.Equal(t, numOfPlayers, game.Round.Turns, "Round should have turns equal to the number of players")
	assert.False(t, game.Round.IsGandalf, "Round should not start with Gandalf called")
}

func TestGetPlayer(t *testing.T) {
	numOfPlayers := 1
	game := InitGame(numOfPlayers)
	mockUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	DefaultUUIDGenerator = func() uuid.UUID {
		return mockUUID
	}
	expectedPlayer := Player{
		Name: "Kazal",
		Id:   mockUUID,
	}

	defer func() { DefaultUUIDGenerator = uuid.New }()

	player, _ := game.GetPlayer(mockUUID)

	assert.Equal(t, expectedPlayer, player, "Should return back player for game")

}
