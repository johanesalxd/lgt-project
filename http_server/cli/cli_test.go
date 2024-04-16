package cli_test

import (
	"strings"
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
	t.Run("record win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &StubPlayerStore{}
		cli := cli.NewCLI(store, in)

		cli.PlayPoker()

		assertPlayerWin(t, store, "Chris")
	})
	t.Run("record win from other user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := &StubPlayerStore{}
		cli := cli.NewCLI(store, in)

		cli.PlayPoker()

		assertPlayerWin(t, store, "Cleo")
	})
}
