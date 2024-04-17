package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.output, PlayerPrompt)
	numberOfPlayers, _ := strconv.Atoi(c.readLine())
	c.scheduleBlindAlerts(numberOfPlayers)

	input := c.readLine()
	c.store.RecordWin(extractWinner(input))
}

func (c *CLI) readLine() string {
	c.input.Scan()

	return c.input.Text()
}

func (c *CLI) scheduleBlindAlerts(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		c.alerter.ScheduledAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
