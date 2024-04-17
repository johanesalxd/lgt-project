package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/johanesalxd/lgt-project/http_server/server"
)

type BlindAlerter interface {
	ScheduledAlertAt(dur time.Duration, amt int)
}

type BlindAlerterFunc func(dur time.Duration, amt int)

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type TexasHoldem struct {
	alerter BlindAlerter
	store   server.PlayerStore
}

type CLI struct {
	input  *bufio.Scanner
	output io.Writer
	game   Game
}

func NewCLI(input io.Reader, output io.Writer, game Game) *CLI {
	return &CLI{
		input:  bufio.NewScanner(input),
		output: output,
		game:   game,
	}
}

func NewTexasHoldem(alerter BlindAlerter, store server.PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}

func StdOutAlerter(dur time.Duration, amt int) {
	time.AfterFunc(dur, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amt)
	})
}
