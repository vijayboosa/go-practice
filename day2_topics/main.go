package main

import (
	"fmt"
	"regexp"
	"strings"
)

func getSomeCordinates() (x, y int) {
	return 12, 12
}

func factorial(x int) int {
	if x <= 0 {
		return 1
	}
	return x * factorial(x-1)
}

func fib(n int) int {
	total := 0
	for l := 1; l <= n; l++ {
		total += l
	}
	return total
}

func tasks() {

	x, y := getSomeCordinates()
	fmt.Println("x:", x, "y:", y)

	fmt.Println("factorial of 4", factorial(4))
	fmt.Println("fib of 10: ", fib(10))

	// reversing a array
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var p1, p2 = 0, len(arr) - 1
	for p1 < p2 {
		arr[p1], arr[p2] = arr[p2], arr[p1]
		p1++
		p2--
	}

	fmt.Println("reversed arr:", arr)

	// filtering even numbers from a slice
	numSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	filteredSlice := []int{}
	for _, x := range numSlice {
		if x%2 == 0 {
			filteredSlice = append(filteredSlice, x)
		}
	}

	fmt.Println("filteredSlice: ", filteredSlice)

	// counting word frequency using map
	sampleText := "One of the primary reasons for using sample text is to test readability. In any design, especially in books or websites, text needs to be clear and easy to read. Sample text helps in evaluating font size, line spacing, and color contrast."

	wordMap := make(map[string]int)

	s := strings.ToLower(sampleText)
	re := regexp.MustCompile(`[^a-z0-9\s]+`)
	s = re.ReplaceAllString(s, " ")

	fmt.Println("Replaced str:", s)
	for x := range strings.SplitSeq(s, " ") {
		_, ok := wordMap[x]
		if !ok {
			wordMap[x] = 1
		} else {
			wordMap[x]++
		}
	}
	fmt.Println("wordmap: ", wordMap)
	for key, value := range wordMap {
		fmt.Printf("Word %s: %d\n", key, value)
	}

}

func main() {

	NewTodoTui().Start()

}
