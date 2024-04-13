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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, wait Sleeper) {
	for i := count; i > 0; i-- {
		fmt.Fprintln(out, i)
		wait.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func NewConfigurableSleeper(duration time.Duration, sleep func(time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{
		duration: duration,
		sleep: sleep,
	}
}
