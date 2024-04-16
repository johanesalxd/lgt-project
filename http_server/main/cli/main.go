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

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := store.NewFSStore(db)
	if err != nil {
		log.Fatalf("can't write store to file %s with error %v", db.Name(), err)
	}

	game := cli.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
