package store_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/johanesalxd/lgt-project/http_server/model"
)

func assertLeague(t testing.TB, got, want model.League) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func createTempFile(t testing.TB, initData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("couldn't create temp file %v", err)
	}

	tmpFile.Write([]byte(initData))

	rmFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, rmFile
}
