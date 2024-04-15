package store

import (
	"encoding/json"
	"io"
	"sync"

	"github.com/johanesalxd/lgt-project/http_server/server"
)

type InMemoryPlayerStore struct {
	store map[string]int
	lock  sync.RWMutex
}

type FSStore struct {
	db io.Reader
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}

func NewFSStore(db io.Reader) *FSStore {
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

func (i *InMemoryPlayerStore) GetLeague() []server.Player {
	var table []server.Player

	for name, wins := range i.store {
		table = append(table, server.Player{Name: name, Wins: wins})
	}

	return table
}

func (f *FSStore) GetLeague() []server.Player {
	var table []server.Player

	json.NewDecoder(f.db).Decode(&table)

	return table
}
