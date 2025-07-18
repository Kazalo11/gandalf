package server

import (
	"log"
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
		log.Printf("Error converting id to uuid: %v", err)
		return
	}

	hub, exists := hubMap[hubId]
	if !exists {
		http.Error(w, "Game not found", http.StatusBadRequest)
		log.Printf("Game not found for id: %s", hubId)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Player"
	}

	game := hub.game
	playerId := uuid.New()
	p := createPlayer(game, playerId, name)
	if p == nil {
		http.Error(w, "Failed to create player", http.StatusInternalServerError)
		log.Println("Failed to create player")
		return
	}
	log.Printf("Created player: %+v", p)

	connectToHub(hub, p, w, r)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	gameId := uuid.New()
	game := internals.InitGame(gameId)
	log.Printf("Created game with id: %s", gameId)

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Player"
	}

	playerId := uuid.New()
	p := createPlayer(game, playerId, name)
	if p == nil {
		http.Error(w, "Failed to create player", http.StatusInternalServerError)
		log.Println("Failed to create player")
		return
	}
	log.Printf("Created player: %+v", p)
	game.AddPlayer(*p)

	hub := newHub(game)
	hubMap[gameId] = hub
	connectToHub(hub, p, w, r)
	go hub.run()
}

func createPlayer(game *internals.Game, playerId uuid.UUID, name string) *models.Player {
	var hand []models.Card
	for i := 0; i < 4; i++ {
		card, err := game.Deck.DrawFromDeck()
		if err != nil {
			log.Printf("Not able to draw from deck: %v", err)
			return nil
		}
		hand = append(hand, card)
	}
	return models.NewPlayer(playerId, name, hand)
}
