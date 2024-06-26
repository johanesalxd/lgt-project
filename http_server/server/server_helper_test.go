package server_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/johanesalxd/lgt-project/http_server/model"
	"github.com/johanesalxd/lgt-project/http_server/server"
)

func newGetScorePostWinRequest(name, method string) *http.Request {
	request, _ := http.NewRequest(method, fmt.Sprintf("/players/%s", name), nil)

	return request
}

func newGetPostLeagueRequest(method string) *http.Request {
	request, _ := http.NewRequest(method, "/league", nil)

	return request
}

func getTableFromBody(t testing.TB, body io.Reader) (table model.League) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&table)
	if err != nil {
		t.Fatalf("failed to parse got %q error %v", body, err)
	}

	return
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertLeague(t testing.TB, got, want model.League) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertContentType(t testing.TB, got string) {
	t.Helper()

	if got != server.ContentType {
		t.Errorf("got %v want %v", got, server.ContentType)
	}
}
