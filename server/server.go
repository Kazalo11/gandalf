package server

import (
	"fmt"
	"net/http"
)

func Start() {
	http.HandleFunc("/ws", wsHandler)
	go handleMessages()
	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
