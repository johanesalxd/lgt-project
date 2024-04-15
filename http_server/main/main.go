package main

import (
	"log"
	"net/http"

	"github.com/johanesalxd/lgt-project/http_server/server"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &server.PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":8498", server))
}
