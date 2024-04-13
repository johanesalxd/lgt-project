package main

import (
	"os"
	"time"

	"github.com/johanesalxd/lgt-project/mocking"
)

func main() {
	// fmt.Println(helloworld.Hello("world", ""))

	// depinj.Greet(os.Stdout, "Elodie")
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(depinj.MyGreeterHandler)))

	sleeper := mocking.NewConfigurableSleeper(2*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
