package main

import (
	"runtime/pprof"
	"time"
)

func potentiallyTimeConsumingCall(result chan struct{}) {
	<-time.After(10 * time.Millisecond)
	result <- struct{}{}
}

func main() {
	for i := 0; i < 100; i++ {
		r := make(chan struct{})
		go potentiallyTimeConsumingCall(r)
		select {
		case <-r:
		case <-time.After(5 * time.Millisecond):
		}
	}

	if pprof.Lookup("goroutine").Count() > 50 {
		panic("seems we are leaking goroutines")
	}
	println("all good")
}
