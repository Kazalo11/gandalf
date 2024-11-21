package server

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Kazalo11/gandalf/models"
	"github.com/google/uuid"
	"golang.org/x/time/rate"
)

type GameServer struct {
	logf                    func(f string, v ...interface{})
	subscriberMessageBuffer int
	publishLimiter          *rate.Limiter
	games                   map[uuid.UUID]*models.Game
	gamesMu                 sync.Mutex
	serveMux                http.ServeMux
}

type subscribe struct {
	msgs      chan []byte
	closeSlow func()
}

func initGameServer() *GameServer {
	gs := &GameServer{
		subscriberMessageBuffer: 16,
		logf:                    log.Printf,
		games:                   make(map[uuid.UUID]*models.Game),
		publishLimiter:          rate.NewLimiter(rate.Every(time.Millisecond*100), 8),
	}

	gs.serveMux.HandleFunc("/join", gs.joinGameHandler)
	gs.serveMux.HandleFunc("/leave", gs.leaveGameHandler)
	gs.serveMux.HandleFunc("/publish", gs.publishHandler)

	return gs
}

func (s *GameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.serveMux.ServeHTTP(w, r)
}

func (s *GameServer) joinGameHandler(w http.ResponseWriter, r *http.Request) {

	gameID := r.URL.Query().Get("game")
	if gameID == "" {
		http.Error(w, "Missing game ID", http.StatusBadRequest)
		return
	}

	ID := uuid.MustParse(gameID)

	s.gamesMu.Lock()
	g, exists := s.games[ID]

	if !exists {
		g = models.InitGame(0)
	}

}
