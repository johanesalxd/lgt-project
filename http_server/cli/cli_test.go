package cli_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

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

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduledAlertAt(dur time.Duration, amt int) {
	s.alerts = append(s.alerts, scheduledAlert{dur, amt})
}

var (
	dummySpyAlerter = &SpyBlindAlerter{}
	dummyStdOut     = &bytes.Buffer{}
)

func TestCLI(t *testing.T) {
	t.Run("record win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &StubPlayerStore{}
		cli := cli.NewCLI(store, in, dummyStdOut, dummySpyAlerter)

		cli.PlayPoker()

		assertPlayerWin(t, store, "Chris")
	})
	t.Run("record win from other user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := &StubPlayerStore{}
		cli := cli.NewCLI(store, in, dummyStdOut, dummySpyAlerter)

		cli.PlayPoker()

		assertPlayerWin(t, store, "Cleo")
	})
	t.Run("scheduled printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &StubPlayerStore{}
		alerter := &SpyBlindAlerter{}

		cli := cli.NewCLI(store, in, dummyStdOut, alerter)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(alerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
				}

				got := alerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
	t.Run("prompts to enter number of players", func(t *testing.T) {
		var (
			dummyPlayerStore = &StubPlayerStore{}
			dummyStdIn       = &bytes.Buffer{}
		)

		cli := cli.NewCLI(dummyPlayerStore, dummyStdIn, dummyStdOut, dummySpyAlerter)
		cli.PlayPoker()

		got := dummyStdOut.String()
		want := "Please enter the number of players: "

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}
