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

type Game struct {
	alerter BlindAlerter
	store   server.PlayerStore
}

type CLI struct {
	input  *bufio.Scanner
	output io.Writer
	game   *Game
}

func NewCLI(input io.Reader, output io.Writer, game *Game) *CLI {
	return &CLI{
		input:  bufio.NewScanner(input),
		output: output,
		game:   game,
	}
}

func NewGame(alerter BlindAlerter, store server.PlayerStore) *Game {
	return &Game{
		alerter: alerter,
		store:   store,
	}
}

func StdOutAlerter(dur time.Duration, amt int) {
	time.AfterFunc(dur, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amt)
	})
}
