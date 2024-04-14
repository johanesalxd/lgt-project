package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, getPlayerScore(player))
}

func getPlayerScore(name string) string {
	switch name {
	case "Pepper":
		return "20"
	case "Floyd":
		return "10"
	}

	return ""
}
