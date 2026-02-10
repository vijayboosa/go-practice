package main

import (
	"fmt"
	"time"
)

func timeout() {

	datach := make(chan int)
	defer close(datach)
	timer := time.NewTimer(200 * time.Millisecond)
	defer timer.Stop()

	go func() {
		for range 5 {
			timer.Reset(200 * time.Millisecond)
			select {
			case d := <-datach:
				fmt.Println("received channel", d)
			case t := <-timer.C:
				fmt.Println("timeout", t)
			}
		}
	}()

	for x := range 5 {
		datach <- x
	}

}
