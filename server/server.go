package server

import (
	"fmt"
	"github.com/rs/cors"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws/create", CreateGame)
	mux.HandleFunc("/ws/game/{id}/join", JoinGame)
	mux.HandleFunc("/game/{id}/player/{playerId}", GetPlayer)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}).Handler(mux)

	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
