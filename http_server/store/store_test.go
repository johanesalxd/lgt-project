package store_test

import (
	"testing"

	"github.com/johanesalxd/lgt-project/http_server/model"
	"github.com/johanesalxd/lgt-project/http_server/store"
)

func TestFSStore(t *testing.T) {
	t.Run("get league table from reader", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store := store.NewFSStore(db)

		got := store.GetLeague()
		want := []model.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()

		assertLeague(t, got, want)
	})
	t.Run("get player score", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store := store.NewFSStore(db)

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})
	t.Run("store player score", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store := store.NewFSStore(db)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		assertScoreEquals(t, got, want)
	})
}
