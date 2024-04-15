package server

import (
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store  PlayerStore
	router *http.ServeMux
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

func NewPlayerServer(store PlayerStore) PlayerServer {
	p := PlayerServer{
		store:  store,
		router: http.NewServeMux(),
	}

	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	p.router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	return p
}
