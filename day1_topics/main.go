package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("hello there")

	currentTime := time.Now()
	fmt.Println("Current time", currentTime)

	// declaring variables
	var i int                   // without initial value
	var name string = "Vijay"   // with value and type
	var mail = "vijay@mail.com" // omitting type (with type inference)

	// var i, j, k int -> this will create a variable without any initial Value
	// var i, j, k int = 1,2,3 -> this will create a varaible with same type
	// var name, mail, isEmp = "Vijay", "vijay@mail.com", true -> using type inference

	fmt.Println("this is to test print statement", i)
	fmt.Printf("Hi, this is %s, and here is my mail %s\n", name, mail)

	// constants and iota
	const Pi = 3.14159                     // untyped numeric constant
	const Greeting string = "good morning" // type string constant

	const (
		StatusOk      = 200
		StatusCreated = 201
		StatusError   = 500
	)

	// we can also use iota which creates a incremental values at compile time
	const (
		Sunday    int = iota // Sunday = 0
		Monday               // Monday = 1
		Tuesday              // Tuesday = 2
		Wednesday            // Wednesday = 3
		Thursday             // Thursday = 4
		Friday               // Friday = 5
		Saturday             // Saturday = 6
	)

	// we can skip a value is iota using _

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
	)
	fmt.Println("int value:", MB)

	const (
		FlagOne = 1 << iota // 1 (0001)
		FlagTwo
		FlagThree
	)

	// type conversion
	vi1 := 23
	vf2 := float64(vi1)
	fmt.Printf("type converion from value %[1]d(%[1]T) -> %[2]f(%[2]T)\n", vi1, vf2)

	vs3 := "232"
	vi4, err := strconv.Atoi(vs3)
	if err != nil {
		fmt.Println("Failed to convert string to int", err)
	}
	fmt.Printf("type conversion from %[1]T -> %[2]T\n", vs3, vi4)

	var value1 int
	var value2 int
	fmt.Print("enter a number: ")
	fmt.Scanln(&value1)
	fmt.Print("enter another number: ")
	fmt.Scanln(&value2)
	fmt.Printf("addition: %d + %d -> %d\n", value1, value2, value1+value2)
}
