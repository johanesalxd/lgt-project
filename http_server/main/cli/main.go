package main

import (
	"fmt"
	"log"
	"os"

	"github.com/johanesalxd/lgt-project/http_server/cli"
	"github.com/johanesalxd/lgt-project/http_server/store"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	store, close, err := store.NewFSStore(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := cli.NewCLI(store, os.Stdin, cli.BlindAlerterFunc(cli.StdOutAlerter))
	game.PlayPoker()
}
