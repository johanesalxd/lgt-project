package store

import (
	"encoding/json"
	"io"
	"os"

	"github.com/johanesalxd/lgt-project/http_server/model"
)

type FSStore struct {
	db     io.Writer
	league model.League
}

type Tape struct {
	file *os.File
}

func NewFSStore(db *os.File) *FSStore {
	db.Seek(0, io.SeekStart)
	league, _ := newTable(db)

	return &FSStore{db: &Tape{db}, league: league}
}

func NewTape(db *os.File) io.Writer {
	return &Tape{file: db}
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

	json.NewEncoder(f.db).Encode(f.league)
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, io.SeekStart)

	return t.file.Write(p)
}
