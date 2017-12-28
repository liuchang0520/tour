package main

import (
	"runtime/pprof"
	"time"
)

const (
	// You shouldn't change const
	waitTime      = 5 * time.Millisecond
	consumingTime = 500 * time.Millisecond
)

func potentiallyTimeConsumingCall(result chan struct{}) {
	<-time.After(consumingTime)
	result <- struct{}{}
}

func main() {
	for i := 0; i < 100; i++ {
		r := make(chan struct{})
		go potentiallyTimeConsumingCall(r)
		select {
		case <-r:
		case <-time.After(waitTime):
		}
	}

	if pprof.Lookup("goroutine").Count() > 50 {
		panic("seems we are leaking goroutines")
	}
	println("all good")
}
