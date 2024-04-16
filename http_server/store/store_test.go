package store_test

import (
	"io"
	"testing"

	"github.com/johanesalxd/lgt-project/http_server/model"
	"github.com/johanesalxd/lgt-project/http_server/store"
)

func TestFSStore(t *testing.T) {
	t.Run("get sorted league table from reader", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()
		store, err := store.NewFSStore(db)

		assertNoError(t, err)

		got := store.GetLeague()
		want := model.League{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
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
		store, err := store.NewFSStore(db)

		assertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})
	t.Run("store player score", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()
		store, err := store.NewFSStore(db)

		assertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		assertScoreEquals(t, got, want)
	})
	t.Run("store new player score", func(t *testing.T) {
		db, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()
		store, err := store.NewFSStore(db)

		assertNoError(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		assertScoreEquals(t, got, want)
	})
	t.Run("test with empty file", func(t *testing.T) {
		file, clean := createTempFile(t, "")
		defer clean()

		_, err := store.NewFSStore(file)

		assertNoError(t, err)
	})
}

func TestTapeWrite(t *testing.T) {
	t.Run("test with existing file", func(t *testing.T) {
		file, clean := createTempFile(t, "12345")
		defer clean()

		tape := store.NewTape(file)
		tape.Write([]byte("abc"))

		file.Seek(0, io.SeekStart)
		newFileContents, _ := io.ReadAll(file)

		got := string(newFileContents)
		want := "abc"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
