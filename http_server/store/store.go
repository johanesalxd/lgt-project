package store

import (
	"encoding/json"
	"io"

	"github.com/johanesalxd/lgt-project/http_server/model"
)

type FSStore struct {
	db     io.ReadWriteSeeker
	league model.League
}

func NewFSStore(db io.ReadWriteSeeker) *FSStore {
	db.Seek(0, io.SeekStart)
	league, _ := newTable(db)

	return &FSStore{db: db, league: league}
}

func (f *FSStore) GetLeague() model.League {
	return f.league
}

func (f *FSStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FSStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, model.Player{Name: name, Wins: 1})
	}

	f.db.Seek(0, io.SeekStart)
	json.NewEncoder(f.db).Encode(f.league)
}
