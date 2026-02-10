package main

import (
	"fmt"
	"sync"
	"time"
)

func buf_ubuf() {

	fmt.Println("strated buf and unbuf channel")
	var wg sync.WaitGroup
	msgCh := make(chan int, 1)

	wg.Go(func() {
		defer close(msgCh)
		for range 10 {
			fmt.Println("before send")
			msgCh <- 1
			fmt.Println("After send")
		}
	})

	wg.Go(func() {
		for range msgCh {
			time.Sleep(1 * time.Second)
			fmt.Println("recevied from channel")
		}
	})

	wg.Wait()

	fmt.Println("Completed buf and unbuf channel")

}
