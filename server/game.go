package server

import (
	"fmt"
	"net/http"

	"github.com/Kazalo11/gandalf/internals"
	"github.com/Kazalo11/gandalf/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	games     = map[uuid.UUID]*internals.Game{}
	gameConns = map[uuid.UUID][]*websocket.Conn{}
)

func JoinGame(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()

	id := r.PathValue("id")

	gameId, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Error converting id to uuid", err)
	}

	game, exists := games[gameId]
	if !exists {
		fmt.Println("Game does not exist, disconnecting")
		conn.Close()
	}

	mutex.Lock()
	clients = append(clients, conn)
	gameConns[gameId] = append(gameConns[gameId], conn)
	mutex.Unlock()

	player := createPlayer(*game)
	if player == nil {
		fmt.Println("Failed to create player")
		conn.Close()
	} else {
		game.AddPlayer(*player)
		message := "Added player to the game"
		broadcast <- []byte(message)
	}

	go broadcastReceivedMessages()
	sendMessage(conn)

}

func createPlayer(game internals.Game) *models.Player {
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

func CreateGame(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}

	mutex.Lock()
	game := internals.InitGame()
	gameId := uuid.New()
	games[gameId] = game

	clients = append(clients, conn)
	gameConns[gameId] = append(gameConns[gameId], conn)
	mutex.Unlock()

	err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Created game with id: %s", gameId)))
	if err != nil {
		fmt.Println("Error writing game Id message: %w", err)
		conn.Close()
	}

	player := createPlayer(*game)
	if player == nil {
		fmt.Println("Failed to create player")
		conn.Close()
	} else {
		game.AddPlayer(*player)
		fmt.Println("Added player to the game")
	}

	go broadcastReceivedMessages()

	sendMessage(conn)

}
