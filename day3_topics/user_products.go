package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type User struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

func (u User) DisplayName() string {
	return strings.TrimSpace(u.Name)
}

type Product struct {
	Title       string    `json:"title"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	InStock     bool      `json:"in_stock"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
}

func (p *Product) Restock(q int) {
	if q <= 0 {
		return
	}

	p.Quantity += q
	p.InStock = p.Quantity > 0
}

// interface which is implemented by User and Product struct
type JSONMarshal interface {
	ToJson() (string, error)
}

func (u User) ToJson() (string, error) {

	b, err := json.MarshalIndent(u, "", " ")

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (p Product) ToJson() (string, error) {

	b, err := json.MarshalIndent(p, "", " ")

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func PrintToJson(j JSONMarshal) {
	s, err := j.ToJson()
	if err != nil {
		fmt.Println("failed to marshal json:", err)
		return
	}

	fmt.Println(s)
}
