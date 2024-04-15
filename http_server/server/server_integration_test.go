package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/johanesalxd/lgt-project/http_server/server"
	"github.com/johanesalxd/lgt-project/http_server/store"
)

func TestRecordAndRetrieveWins(t *testing.T) {
	svr := server.NewPlayerServer(store.NewInMemoryPlayerStore())
	player := "Pepper"

	response := httptest.NewRecorder()
	request := newGetScorePostWinRequest(player, http.MethodPost)

	svr.ServeHTTP(response, request)
	svr.ServeHTTP(response, request)
	svr.ServeHTTP(response, request)

	t.Run("get score", func(t *testing.T) {
		response = httptest.NewRecorder()
		request = newGetScorePostWinRequest(player, http.MethodGet)

		svr.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})
	t.Run("get league table", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := newGetPostLeagueRequest(http.MethodGet)

		svr.ServeHTTP(response, request)

		got := getTableFromBody(t, response.Body)
		want := []server.Player{
			{
				Name: "Pepper",
				Wins: 3,
			},
		}
		assertLeague(t, got, want)
	})
}
