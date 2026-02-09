package main

import (
	"fmt"
	"strings"
	"time"
)

type todo struct {
	id      int
	title   string
	created time.Time
	status  bool
}

type TodoTui struct {
	Todos []todo
}

func NewTodoTui() *TodoTui {
	return &TodoTui{}
}

func (t *TodoTui) addTask() {
	var task string
	for {
		fmt.Print("Enter task title:")
		fmt.Scanln(&task)

		task = strings.TrimSpace(task)

		if len(task) == 0 {
			continue
		}

		t.Todos = append(t.Todos, todo{
			id:      len(t.Todos) + 1,
			title:   task,
			created: time.Now(),
			status:  false,
		})
		break
	}

}

func (t *TodoTui) listTasks() {
	var option string
	for {

		fmt.Println("List:")
		fmt.Println("1> All")
		fmt.Println("2> Pending")
		fmt.Println("3> Done")
		fmt.Print("Choose:")
		fmt.Scanln(&option)

		switch option {
		case "1":
			fmt.Println("--- All tasks ---")
			for _, todo := range t.Todos {
				fmt.Printf("[TODO] #%d %s")
			}
		case "2":
			fmt.Println("option choosed")
		case "3":
			fmt.Println("option choosed")
		}

	}

}

func (t *TodoTui) initalMenu() {
	var selectedOption string
	fmt.Println("======= TO-DO CLI =======")
	fmt.Println("1> Add task")
	fmt.Println("2> List tasks")
	fmt.Println("3> Mark task as done")
	fmt.Println("4> Delete task")
	fmt.Println("5> Search task")
	fmt.Println("6>Exit")
	for {
		fmt.Print("Choose:")
		fmt.Scanln(&selectedOption)

		switch selectedOption {
		case "1":
			fmt.Println("option choosed")
		case "2":
			fmt.Println("option choosed")
		case "3":
			fmt.Println("option choosed")
		case "4":
			fmt.Println("option choosed")

		case "5":
			fmt.Println("option choosed")

		case "6":
			fmt.Println("option choosed")

		default:
			fmt.Println("Invalid choice. Please select 1-6")
		}

	}

}

func (t *TodoTui) Start() {

}
