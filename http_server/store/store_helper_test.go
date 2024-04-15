package store_test

import (
	"reflect"
	"testing"

	"github.com/johanesalxd/lgt-project/http_server/server"
)

func assertLeague(t testing.TB, got, want []server.Player) {
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