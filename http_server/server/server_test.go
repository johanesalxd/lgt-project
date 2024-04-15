package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/johanesalxd/lgt-project/http_server/server"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []server.Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []server.Player {
	return s.league
}

func TestGETPlayer(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := server.NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScorePostWinRequest("Pepper", http.MethodGet)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScorePostWinRequest("Floyd", http.MethodGet)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})
	t.Run("return 404 on missing players", func(t *testing.T) {
		request := newGetScorePostWinRequest("Apollo", http.MethodGet)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{},
	}
	server := server.NewPlayerServer(&store)

	t.Run("record wins on POST and returns 202", func(t *testing.T) {
		player := "Pepper"
		request := newGetScorePostWinRequest(player, http.MethodPost)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	svr := server.NewPlayerServer(&store)

	t.Run("return 200 on league", func(t *testing.T) {
		request := newGetPostLeagueRequest(http.MethodGet)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		var got []server.Player

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("failed to parse got %q error %v", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("returns league table as JSON", func(t *testing.T) {
		table := []server.Player{
			{Name: "Cleo",
				Wins: 32},
			{Name: "Chris",
				Wins: 20},
			{Name: "Tiest",
				Wins: 14},
		}

		store := StubPlayerStore{league: table}
		svr := server.NewPlayerServer(&store)

		request := newGetPostLeagueRequest(http.MethodGet)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		got := assertLeagueResponseBody(t, response.Body)
		assertLeague(t, got, table)
	})
}
