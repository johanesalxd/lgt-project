package main

import (
	"log"
	"net/http"
	"os"

	"github.com/johanesalxd/lgt-project/http_server/server"
	"github.com/johanesalxd/lgt-project/http_server/store"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	server := server.NewPlayerServer(store.NewFSStore(db))
	log.Fatal(http.ListenAndServe(":8498", server))
}
