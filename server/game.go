package server

import (
	"log"
	"net/http"

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
	if len(game.Players) == 4 {
		http.Error(w, "Game is full", http.StatusBadRequest)
		log.Println("Game is full")
		return
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

	connectToHub(hub, p, w, r)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	gameId := uuid.New()
	game := models.InitGame(gameId)
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

func ReconnectGame(w http.ResponseWriter, r *http.Request) {
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

	playerId := r.URL.Query().Get("playerId")
	if playerId == "" {
		http.Error(w, "Player ID is required", http.StatusBadRequest)
		log.Println("Player ID is required")
		return
	}

	playerUUID, err := uuid.Parse(playerId)
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		log.Printf("Error converting playerId to uuid: %v", err)
		return
	}

	p, exists := hub.game.Players[playerUUID]
	if !exists {
		http.Error(w, "Player not found", http.StatusNotFound)
		log.Printf("Player not found for id: %s", playerId)
		return
	}

	connectToHub(hub, p, w, r)
}
