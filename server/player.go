package server

import (
	"encoding/json"
	"github.com/Kazalo11/gandalf/models"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	playerId := r.PathValue("playerId")

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

	playerUUID, err := uuid.Parse(playerId)
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		log.Printf("Error converting playerId to uuid: %v", err)
		return
	}

	player, exists := hub.game.Players[playerUUID]
	if !exists {
		http.Error(w, "Player not found", http.StatusNotFound)
		log.Printf("Player not found for id: %s", playerId)
		return
	}

	response, err := json.Marshal(player)
	if err != nil {
		http.Error(w, "Failed to marshal player data", http.StatusInternalServerError)
		log.Printf("Error marshalling player data: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func createPlayer(game *models.Game, playerId uuid.UUID, name string) *models.Player {
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
