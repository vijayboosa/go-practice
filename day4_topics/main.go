package main

import (
	"fmt"
	"sync"
)

func swap(a, b **int) {
	*a, *b = *b, *a
}

type User struct {
	Name string
	Mail string
}

func (u *User) UpdateName(n string) {
	u.Name = n
}

func main() {
	x := 10
	y := 29

	px := &x
	py := &y

	// swaping pointers
	fmt.Printf("Before swaping: px=%p -> %d | py=%p -> %d\n", px, *px, py, *py)

	swap(&px, &py)

	fmt.Printf("After swaping: px=%p -> %d | py=%p -> %d\n", px, *px, py, *py)

	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		i := i
		wg.Go(func() {
			for j := range 10 {
				fmt.Printf("grc #%d -> value %d\n", i, j)
			}
		})
	}

	wg.Wait()
	fmt.Println("All groutines are completed")
}
