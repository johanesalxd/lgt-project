package cli

import (
	"bufio"
	"strings"
)

func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.input)
	reader.Scan()
	c.store.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
