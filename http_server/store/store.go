package store

import (
	"encoding/json"
	"io"
	"sync"

	"github.com/johanesalxd/lgt-project/http_server/model"
)

type InMemoryPlayerStore struct {
	store map[string]int
	lock  sync.RWMutex
}

type FSStore struct {
	db io.ReadWriteSeeker
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}

func NewFSStore(db io.ReadWriteSeeker) *FSStore {
	return &FSStore{db: db}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.RLock()
	defer i.lock.RUnlock()

	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() model.League {
	var table model.League

	for name, wins := range i.store {
		table = append(table, model.Player{Name: name, Wins: wins})
	}

	return table
}

func (f *FSStore) GetLeague() model.League {
	f.db.Seek(0, io.SeekStart)

	table, _ := f.newTable(f.db)

	return table
}

func (f *FSStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FSStore) RecordWin(name string) {
	table := f.GetLeague()
	player := table.Find(name)

	if player != nil {
		player.Wins++
	}

	f.db.Seek(0, io.SeekStart)
	json.NewEncoder(f.db).Encode(table)
}
