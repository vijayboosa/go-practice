package main

import (
	"fmt"
	"math/rand"
	"time"
)

func raceChannels() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	w1, w2 := 0, 0
	for range 10 {
		go func() {
			delay := time.Duration(10+rand.Intn(191)) * time.Millisecond
			time.Sleep(delay)
			ch1 <- struct{}{}
		}()

		go func() {
			delay := time.Duration(10+rand.Intn(191)) * time.Millisecond
			time.Sleep(delay)
			ch2 <- struct{}{}
		}()

		select {
		case <-ch1:
			w1++
		case <-ch2:
			w2++
		}

	}

	fmt.Printf("Channel 1: %d wins\n", w1)
	fmt.Printf("Channel 2: %d wins\n", w2)

}
