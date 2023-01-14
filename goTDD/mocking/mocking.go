package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Advanced example of TDD applied on mocked tests

const (
	countingStart = 3
	finalPrint    = "Go!"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{duration: 3 * time.Second, sleep: time.Sleep}
	CountDown(os.Stdout, sleeper)
}

func CountDown(w io.Writer, sleeper Sleeper) {
	for i := countingStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
	sleeper.Sleep()
	fmt.Fprint(w, finalPrint)
}
