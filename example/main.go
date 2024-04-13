package main

import (
	"os"

	"github.com/johanesalxd/lgt-project/mocking"
)

func main() {
	// fmt.Println(helloworld.Hello("world", ""))

	// depinj.Greet(os.Stdout, "Elodie")
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(depinj.MyGreeterHandler)))

	mocking.Countdown(os.Stdout)
}
