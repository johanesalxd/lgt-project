package cli_test

import (
	"testing"

	"github.com/johanesalxd/lgt-project/http_server/cli"
	"github.com/johanesalxd/lgt-project/http_server/model"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   model.League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() model.League {
	return s.league
}

func TestCLI(t *testing.T) {
	store := &StubPlayerStore{}
	cli := cli.NewCLI(store)
	cli.PlayPoker()

	if len(store.winCalls) != 1 {
		t.Fatalf("expected a win call but didn't get any")
	}
}
