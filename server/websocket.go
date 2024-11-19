package server

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/coder/websocket"
	"golang.org/x/time/rate"
)

type messageServer struct {
	logf                    func(f string, v ...interface{})
	subscriberMessageBuffer int
	publishLimiter          *rate.Limiter
	subscribers             map[*subscriber]struct{}
	serveMux                http.ServeMux
	subscribersMu           sync.Mutex
}
type subscriber struct {
	msgs      chan []byte
	closeSlow func()
}

func initMessageServer() *messageServer {
	ms := &messageServer{
		subscriberMessageBuffer: 16,
		logf:                    log.Printf,
		subscribers:             make(map[*subscriber]struct{}),
		publishLimiter:          rate.NewLimiter(rate.Every(time.Millisecond*100), 8),
	}

	ms.serveMux.HandleFunc("/subscribe", ms.subscribeHandler)
	ms.serveMux.HandleFunc("/publish", ms.publishHandler)
}

func (s *messageServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.serveMux.ServeHTTP(w, r)

}

func writeTimeout(ctx context.Context, timeout time.Duration, c *websocket.Conn, msg []byte) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Write(ctx, websocket.MessageText, msg)
}
