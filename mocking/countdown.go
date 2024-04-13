package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord = "Go!"
	count     = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, wait Sleeper) {
	for i := count; i > 0; i-- {
		fmt.Fprintln(out, i)
		wait.Sleep()
	}

	fmt.Fprint(out, finalWord)
}
