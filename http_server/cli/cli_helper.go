package cli

import (
	"strings"
	"time"
)

func (c *CLI) PlayPoker() {
	input := c.readLine()

	c.store.RecordWin(extractWinner(input))
	c.scheduleBlindAlerts()
}

func (c *CLI) readLine() string {
	c.input.Scan()

	return c.input.Text()
}

func (c *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		c.alerter.ScheduledAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
