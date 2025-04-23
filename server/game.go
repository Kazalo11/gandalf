package server

import (
	"fmt"
	"net/http"

	"github.com/Kazalo11/gandalf/internals"
	"github.com/Kazalo11/gandalf/models"
	"github.com/google/uuid"
)

func JoinGame(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	hubId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid game ID", http.StatusBadRequest)
		fmt.Println("Error converting id to uuid", err)
		return
	}

	hub, exists := hubMap[hubId]
	if !exists {
		http.Error(w, "Game not found", http.StatusBadRequest)
		fmt.Println("Game not found for id:", hubId)
		return
	}

	connectToHub(hub, w, r)

}

func CreateGame(w http.ResponseWriter, r *http.Request) {

	hubId := uuid.New()

	game := internals.InitGame()
	fmt.Printf("Created game with id: %s\n", hubId)

	hub := newHub(game)

	hubMap[hubId] = hub
	connectToHub(hub, w, r)

	go hub.run()

}

func createPlayer(game *internals.Game) *models.Player {
	playerId := uuid.New()
	name := "Kazal"
	hand := []models.Card{}
	for i := 0; i < 4; i++ {
		card, err := game.Deck.DrawFromDeck()
		if err != nil {
			fmt.Println("Not able to draw from deck:", err)
			return nil
		}
		hand = append(hand, card)
	}

	return models.NewPlayer(playerId, name, hand)

}
