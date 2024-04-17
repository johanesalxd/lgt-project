package main

import (
	"log"
	"net/http"

	"github.com/johanesalxd/lgt-project/http_server/server"
	"github.com/johanesalxd/lgt-project/http_server/store"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := store.NewFSStore(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := server.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":8498", server))
}
