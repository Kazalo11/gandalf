package main

import (
	"time"

	"github.com/Kazalo11/gandalf/server"
	"golang.org/x/exp/rand"
)

func main() {
	rand.Seed(uint64(time.Now().UnixNano()))
	server.Start()
}
