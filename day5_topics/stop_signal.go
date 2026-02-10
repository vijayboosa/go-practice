package main

import (
	"fmt"
	"sync"
	"time"
)

func stopSignal() {

	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Go(func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		for {

			select {
			case <-stop:
				fmt.Println("bye")
				return

			case <-ticker.C:
				fmt.Println("Doing some work")
			}
		}
	})

	time.Sleep(700 * time.Millisecond)
	close(stop)
	// time.Sleep(50 * time.Millisecond)
}
