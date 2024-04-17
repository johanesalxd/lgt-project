package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

func (b BlindAlerterFunc) ScheduledAlertAt(dur time.Duration, amt int) {
	b(dur, amt)
}

func (g *Game) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		g.alerter.ScheduledAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (g *Game) Finish(winner string) {
	g.store.RecordWin(winner)
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.output, PlayerPrompt)

	numberOfPlayersInput := c.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	c.game.Start(numberOfPlayers)

	winnerInput := c.readLine()
	winner := extractWinner(winnerInput)

	c.game.Finish(winner)
}

func (c *CLI) readLine() string {
	c.input.Scan()

	return c.input.Text()
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
