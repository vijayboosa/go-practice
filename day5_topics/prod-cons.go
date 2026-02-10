package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(j chan int) {
	defer close(j)
	for x := range 50 {
		j <- x
	}
}

func worker(j <-chan int, r chan<- int, wid int) {
	for work := range j {
		fmt.Printf("Worker %d: processing %d\n", wid, work)
		slpTime := (work % 5) * 100
		time.Sleep(time.Duration(slpTime * int(time.Millisecond)))
		fmt.Printf("Worker %d: jobId %d Done\n", wid, work)
		r <- work
	}
}

func consumer(wg *sync.WaitGroup, j <-chan int, r chan<- int, w int) {
	defer close(r) // channel should be closed only by the writers
	for i := range w {
		wg.Go(func() {
			worker(j, r, i)
		})
	}

	wg.Wait()

}

func prodCons() {
	fmt.Println("Starting producer and consumer")
	jobsch := make(chan int)
	resultch := make(chan int)
	result := make([]int, 51)

	var wg sync.WaitGroup

	go producer(jobsch)

	// consumer is blocking
	go consumer(&wg, jobsch, resultch, 4)

	for r := range resultch {
		result[r] = r
	}

	fmt.Printf("job result slice: %v\n", result)
	fmt.Println("Completed producer and consumer")
}
