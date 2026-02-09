package main

import (
	"fmt"
	"time"
)

func reverseUnicode(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	// interating over unicode string
	greet := "1i1"
	for index, x := range greet {
		fmt.Printf("Byte index: %d, Rune: %[2]c, Unicode Code point: %#[2]U\n", index, x)
	}

	s := reverseUnicode(greet)
	fmt.Println("reversed string", s)

	// check if string is palindromne
	if s == greet {
		fmt.Println("found palindromne", s, greet)
	}
	u := User{Name: "  Vijay  ", Mail: "vijay@example.com"}
	p := Product{
		Title:       "PLA Filament",
		Price:       1299.0,
		Quantity:    0,
		InStock:     false,
		CreatedAt:   time.Now(),
		Description: "1kg spool",
	}

	fmt.Println("DisplayName:", u.DisplayName())

	p.Restock(5)

	PrintToJson(u)
	PrintToJson(p)
}
