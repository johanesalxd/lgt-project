package server_test

import (
	"fmt"
	"net/http"
	"testing"
)

func newGetScorePostWinRequest(name, method string) *http.Request {
	request, _ := http.NewRequest(method, fmt.Sprintf("/players/%s", name), nil)

	return request
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
