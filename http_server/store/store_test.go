package store_test

import (
	"strings"
	"testing"

	"github.com/johanesalxd/lgt-project/http_server/server"
	"github.com/johanesalxd/lgt-project/http_server/store"
)

func TestFSStore(t *testing.T) {
	t.Run("get league table from reader", func(t *testing.T) {
		db := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		store := store.NewFSStore(db)

		got := store.GetLeague()
		want := []server.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()

		assertLeague(t, got, want)
	})
	t.Run("get player score", func(t *testing.T) {
		db := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		store := store.NewFSStore(db)

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})
}
