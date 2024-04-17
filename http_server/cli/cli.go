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

func (b BlindAlerterFunc) ScheduledAlertAt(dur time.Duration, amt int) {
	b(dur, amt)
}

type CLI struct {
	store   server.PlayerStore
	input   *bufio.Scanner
	output  io.Writer
	alerter BlindAlerter
}

func NewCLI(store server.PlayerStore, input io.Reader, output io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		store:   store,
		input:   bufio.NewScanner(input),
		output:  output,
		alerter: alerter,
	}
}

func StdOutAlerter(dur time.Duration, amt int) {
	time.AfterFunc(dur, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amt)
	})
}
