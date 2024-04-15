package main

import (
	"log"
	"net/http"

	"github.com/johanesalxd/lgt-project/http_server/server"
	"github.com/johanesalxd/lgt-project/http_server/store"
)

func main() {
	server := &server.PlayerServer{Store: store.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":8498", server))
}
