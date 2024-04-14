package main

import (
	"log"
	"net/http"

	"github.com/johanesalxd/lgt-project/http_server/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":8498", handler))
}
