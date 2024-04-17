package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/johanesalxd/lgt-project/http_server/model"
	"github.com/johanesalxd/lgt-project/http_server/server"
	"github.com/johanesalxd/lgt-project/http_server/store"
)

func TestRecordAndRetrieveWins(t *testing.T) {
	db, cleanDB := createTempFile(t, "[]")
	defer cleanDB()

	store, err := store.NewTestFSStore(db)

	assertNoError(t, err)

	svr := server.NewPlayerServer(store)
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
		want := model.League{
			{
				Name: "Pepper",
				Wins: 3,
			},
		}
		assertLeague(t, got, want)
	})
}
