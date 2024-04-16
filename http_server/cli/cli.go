package cli

import "github.com/johanesalxd/lgt-project/http_server/server"

type CLI struct {
	store server.PlayerStore
}

func NewCLI(store server.PlayerStore) *CLI {
	p := new(CLI)
	p.store = store

	return p
}

func (c *CLI) PlayPoker() {
	c.store.RecordWin("Cleo")
}
