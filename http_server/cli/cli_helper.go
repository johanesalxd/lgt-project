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

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.output, PlayerPrompt)

	numberOfPlayersInput := c.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))
	if err != nil {
		fmt.Fprint(c.output, "you're so silly")
		return
	}

	c.game.Start(numberOfPlayers)

	winnerInput := c.readLine()
	winner := extractWinner(winnerInput)

	c.game.Finish(winner)
}

func (c *CLI) readLine() string {
	c.input.Scan()

	return c.input.Text()
}

func (t *TexasHoldem) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		t.alerter.ScheduledAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (t *TexasHoldem) Finish(winner string) {
	t.store.RecordWin(winner)
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
