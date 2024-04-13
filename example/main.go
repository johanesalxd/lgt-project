package main

import (
	"log"
	"net/http"
	"os"

	depinj "github.com/johanesalxd/lgt-project/dep_inj"
)

func main() {
	// fmt.Println(helloworld.Hello("world", ""))

	depinj.Greet(os.Stdout, "Elodie")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(depinj.MyGreeterHandler)))
}
