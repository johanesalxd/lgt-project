package store

import (
	"sync"

	"github.com/johanesalxd/lgt-project/http_server/model"
)

type InMemoryPlayerStore struct {
	store map[string]int
	lock  sync.RWMutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
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
