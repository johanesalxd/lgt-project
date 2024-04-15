package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/johanesalxd/lgt-project/http_server/server"
	"github.com/johanesalxd/lgt-project/http_server/store"
)

func TestRecordAndRetrieveWins(t *testing.T) {
	server := server.PlayerServer{store.NewInMemoryPlayerStore()}
	player := "Pepper"

	response := httptest.NewRecorder()
	request := newGetScorePostWinRequest(player, http.MethodPost)

	server.ServeHTTP(response, request)
	server.ServeHTTP(response, request)
	server.ServeHTTP(response, request)

	response = httptest.NewRecorder()
	request = newGetScorePostWinRequest(player, http.MethodGet)
	server.ServeHTTP(response, request)

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
