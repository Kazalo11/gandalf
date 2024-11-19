package server

import (
	"context"
	"io"
	"net/http"
)

func (ms *messageServer) publish(msg []byte) {
	ms.subscribersMu.Lock()
	defer ms.subscribersMu.Unlock()

	ms.publishLimiter.Wait(context.Background())

	for s := range ms.subscribers {
		select {
		case s.msgs <- msg:
		default:
			go s.closeSlow()
		}
	}
}

func (ms *messageServer) publishHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	body := http.MaxBytesReader(w, r.Body, 8192)
	msg, err := io.ReadAll(body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusRequestEntityTooLarge), http.StatusRequestEntityTooLarge)
		return
	}

	ms.publish(msg)

	w.WriteHeader(http.StatusAccepted)
}
