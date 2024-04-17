package store

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/johanesalxd/lgt-project/http_server/model"
)

type FSStore struct {
	db     *json.Encoder
	league model.League
}

func NewFSStore(dbFileName string) (*FSStore, func(), error) {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := newFSStore(db)
	if err != nil {
		log.Fatalf("can't write store to file %s with error %v", db.Name(), err)
	}

	return store, closeFunc, nil
}

func NewTestFSStore(db *os.File) (*FSStore, error) {
	return newFSStore(db)
}

type Tape struct {
	file *os.File
}

func NewTape(db *os.File) io.Writer {
	return &Tape{file: db}
}

func (f *FSStore) GetLeague() model.League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
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

	f.db.Encode(f.league)
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, io.SeekStart)

	return t.file.Write(p)
}

func newFSStore(db *os.File) (*FSStore, error) {
	err := initDB(db)
	if err != nil {
		return nil, fmt.Errorf("can't init store with error %v", err)
	}

	league, err := newTable(db)
	if err != nil {
		return nil, fmt.Errorf("can't load store from file %s with error %v", db.Name(), err)
	}

	return &FSStore{
		db:     json.NewEncoder(&Tape{db}),
		league: league,
	}, nil
}
