package main

import (
	"sync"
	"testing"
	"time"
)

func task() {
	time.Sleep(100 * time.Millisecond)
}

func BenchmarkSequential(b *testing.B) {

	for b.Loop() {
		task()
	}

}

func BenchmarkConcurrent(b *testing.B) {

	var wg sync.WaitGroup

	for b.Loop() {
		wg.Go(func() {
			task()

		})
	}

	wg.Wait()

}
