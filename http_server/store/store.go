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

func (i *InMemoryPlayerStore) GetLeague() []model.Player {
	var table []model.Player

	for name, wins := range i.store {
		table = append(table, model.Player{Name: name, Wins: wins})
	}

	return table
}

func (f *FSStore) GetLeague() []model.Player {
	f.db.Seek(0, io.SeekStart)

	table, _ := f.newTable(f.db)

	return table
}

func (f *FSStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins

			break
		}
	}

	return wins
}

func (f *FSStore) RecordWin(name string) {
	table := f.GetLeague()

	for i, player := range table {
		if player.Name == name {
			table[i].Wins++
		}
	}

	f.db.Seek(0, io.SeekStart)
	json.NewEncoder(f.db).Encode(table)
}
