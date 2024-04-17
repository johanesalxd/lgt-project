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

type GameSpy struct {
	startedWith  int
	finishedWith string
	startCalled  bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.startCalled = true
	g.startedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.finishedWith = winner
}

var (
	dummySpyAlerter  = &SpyBlindAlerter{}
	dummyPlayerStore = &StubPlayerStore{}
	dummyStdOut      = &bytes.Buffer{}
)

func TestCLI(t *testing.T) {
	t.Run("scheduled printing of blind values", func(t *testing.T) {
		// in := strings.NewReader("Chris wins\n")
		in := strings.NewReader("5\n")
		store := &StubPlayerStore{}
		alerter := &SpyBlindAlerter{}

		game := cli.NewTexasHoldem(alerter, store)
		cli := cli.NewCLI(in, dummyStdOut, game)
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
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		alerter := &SpyBlindAlerter{}

		game := cli.NewTexasHoldem(alerter, dummyPlayerStore)
		cmd := cli.NewCLI(in, stdout, game)
		cmd.PlayPoker()

		got := stdout.String()
		want := cli.PlayerPrompt

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
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
	t.Run("prompts user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		input := strings.NewReader("7\n")
		game := &GameSpy{}

		cmd := cli.NewCLI(input, stdout, game)
		cmd.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := cli.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q want %q", gotPrompt, wantPrompt)
		}

		if game.startedWith != 7 {
			t.Errorf("wanted Start called out with 7 but got %d", game.startedWith)
		}
	})
	t.Run("error when non numeric value is entered", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		input := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cmd := cli.NewCLI(input, stdout, game)
		cmd.PlayPoker()

		if game.startCalled {
			t.Errorf("game should not have started")
		}

		gotPrompt := stdout.String()
		wantPrompt := cli.PlayerPrompt + "you're so silly"

		if gotPrompt != wantPrompt {
			t.Errorf("got %q want %q", gotPrompt, wantPrompt)
		}
	})
}

func TestGameStart(t *testing.T) {
	t.Run("schedules alert on game start for 5 players", func(t *testing.T) {
		alerter := &SpyBlindAlerter{}
		game := cli.NewTexasHoldem(alerter, dummyPlayerStore)

		game.Start(5)

		cases := []scheduledAlert{
			{at: 0 * time.Second, amount: 100},
			{at: 10 * time.Minute, amount: 200},
			{at: 20 * time.Minute, amount: 300},
			{at: 30 * time.Minute, amount: 400},
			{at: 40 * time.Minute, amount: 500},
			{at: 50 * time.Minute, amount: 600},
			{at: 60 * time.Minute, amount: 800},
			{at: 70 * time.Minute, amount: 1000},
			{at: 80 * time.Minute, amount: 2000},
			{at: 90 * time.Minute, amount: 4000},
			{at: 100 * time.Minute, amount: 8000},
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
	t.Run("schedules alert on game to start for 7 players", func(t *testing.T) {
		alerter := &SpyBlindAlerter{}
		game := cli.NewTexasHoldem(alerter, dummyPlayerStore)

		game.Start(7)

		cases := []scheduledAlert{
			{at: 0 * time.Second, amount: 100},
			{at: 12 * time.Minute, amount: 200},
			{at: 24 * time.Minute, amount: 300},
			{at: 36 * time.Minute, amount: 400},
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
}

func TestGameFinish(t *testing.T) {
	store := &StubPlayerStore{}
	game := cli.NewTexasHoldem(dummySpyAlerter, store)
	winner := "Ruth"

	game.Finish(winner)

	assertPlayerWin(t, store, winner)
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}
