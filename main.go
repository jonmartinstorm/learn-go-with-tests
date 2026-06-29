package main

import (
	"os"
	"time"

	"github.com/jonmartinstorm/learn-go-with-tests/mocking"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleeperFunc: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
