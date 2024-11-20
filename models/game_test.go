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

	mockUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	DefaultUUIDGenerator = func() uuid.UUID {
		return mockUUID
	}
	expectedPlayer := Player{
		Name: "Kazal",
		Id:   mockUUID,
	}
	game := InitGame(numOfPlayers)

	player, _ := game.GetPlayer(mockUUID)

	assert.Equal(t, expectedPlayer.Name, player.Name, "Should return back player for game")

}

func TestEndGameWithoutGandalf(t *testing.T) {
	game := InitGame(2)

	assert.False(t, game.Round.IsGandalf, "Gandalf should not be called at the start")

	ended, winner := game.EndGame()
	assert.False(t, ended, "EndGame should return false if Gandalf was not called")
	assert.Equal(t, Player{}, winner, "EndGame should not return a winner if Gandalf was not called")
}

func TestEndWithGandalf(t *testing.T) {
	game := InitGame(2)
	game.Round.IsGandalf = true
	game.Round.Turns = 2
	game.Round.CurrentTurn = 2
	players := []Player{}

	hand1 := []Card{
		{
			Rank: KING,
			Suit: SPADE,
		},
		{
			Rank: SEVEN,
			Suit: HEART,
		},
	}

	hand2 := []Card{
		{
			Rank: FOUR,
			Suit: SPADE,
		},
		{
			Rank: SIX,
			Suit: HEART,
		},
	}

	player1 := Player{
		Name: "Kazal",
		Hand: hand1,
	}
	player2 := Player{
		Name: "Tom",
		Hand: hand2,
	}
	players = append(players, player1)
	players = append(players, player2)

	game.Players = players

	boolean, player := game.EndGame()

	assert.True(t, boolean, "if isGandalf is true then should end the game")

	assert.Equal(t, player1, player, "Should return the winning player")

}
