package main

import (
	"fmt"
	"sync"
)

func pingPong() {
	fmt.Println("start ping pong")
	pingch := make(chan string)
	pongch := make(chan string)

	var wg sync.WaitGroup

	// GA
	wg.Go(func() {
		defer close(pingch)
		for i := range 5 {
			pingch <- "ping"
			fmt.Printf("ping sent%d\n", i)
		}
	})

	// GB
	wg.Go(func() {
		defer close(pongch)
		for range pingch {
			pongch <- "pong"
		}
	})

	for range pongch {
		fmt.Println("recevied pong")
	}

	wg.Wait()

	fmt.Println("All goroutines are completed")
	fmt.Println("completed ping pong")
}
