package cli

import (
	"io"

	"github.com/johanesalxd/lgt-project/http_server/server"
)

type CLI struct {
	store server.PlayerStore
	input io.Reader
}

func NewCLI(store server.PlayerStore, input io.Reader) *CLI {
	p := new(CLI)

	p.store = store
	p.input = input

	return p
}
