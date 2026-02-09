package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Todo struct {
	id      int
	title   string
	created time.Time
	status  bool
}

type TodoTui struct {
	Todos  map[int]Todo
	nextID int
}

func NewTodoTui() *TodoTui {
	return &TodoTui{
		Todos:  make(map[int]Todo),
		nextID: 1,
	}
}

func (t *TodoTui) addTask() {
	var task string
	for {
		fmt.Print("Enter task title:")

		reader := bufio.NewReader(os.Stdin)
		task, _ = reader.ReadString('\n')
		task = strings.TrimSpace(task)

		if len(task) == 0 {
			fmt.Println("Title cannot be empty.")
			continue
		}

		newTodoId := t.nextID
		t.nextID++
		t.Todos[newTodoId] = Todo{
			id:      newTodoId,
			title:   task,
			created: time.Now(),
			status:  false,
		}
		fmt.Printf("Task added. (ID:%d)\n", newTodoId)
		break
	}

}

func (t *TodoTui) listTasks() {
	var option string
	fmt.Println("List:")
	fmt.Println("1> All")
	fmt.Println("2> Pending")
	fmt.Println("3> Done")
	for {

		fmt.Print("Choose:")
		fmt.Scanln(&option)

		switch option {
		case "1":
			fmt.Println("--- All tasks ---")
			if len(t.Todos) == 0 {
				println("(none)")
				break
			}
			for _, todo := range t.Todos {
				status := "TODO"
				if todo.status {
					status = "DONE"
				}

				fmt.Printf("[%s] #%d %-30s(%s)\n", status, todo.id, todo.title, todo.created.Format("02-01-2006 3:04PM"))
			}
		case "2":
			fmt.Println("--- Pending tasks ---")
			if len(t.Todos) == 0 {
				println("(none)")
				break
			}

			for _, todo := range t.Todos {
				fmt.Printf("[TODO] #%d %-30s%s\n", todo.id, todo.title, todo.created.Format("02-01-2006 3:04PM"))
			}
		case "3":
			fmt.Println("--- Dond tasks ---")
			if len(t.Todos) == 0 {
				println("(none)")
				break
			}
			for id, todo := range t.Todos {
				if todo.status {
					fmt.Printf("[DONE] #%d %-30s%s\n", id, todo.title, todo.created.Format("02-01-2006 3:04PM"))
				}
			}

		default:
			fmt.Println("enter a valid option:")
			continue
		}

		break

	}

}

func (t *TodoTui) markTaskDone() {
	var todoId int
	fmt.Print("Enter task ID to mark done:")

	for {
		_, err := fmt.Scanln(&todoId)

		if err != nil {
			fmt.Println("enter a valid number:")
			continue
		}

		todo, ok := t.Todos[todoId]

		if ok {
			if todo.status {
				fmt.Println("Already done.")
			} else {
				todo.status = true
				t.Todos[todoId] = todo
				fmt.Println("Marked as done.")
			}
		} else {
			fmt.Println("Task ID not found.")
		}
		break
	}
}

func (t *TodoTui) deleteTask() {
	var todoId int
	fmt.Print("Enter task ID to delete:")

	for {
		_, err := fmt.Scanln(&todoId)

		if err != nil {
			fmt.Println("enter a valid number:")
			continue
		}

		_, ok := t.Todos[todoId]

		if ok {
			delete(t.Todos, todoId)
			fmt.Println("Deleted.")
		} else {
			fmt.Println("Task ID not found.")
		}
		break
	}
}

func (t *TodoTui) searchTask() {

	fmt.Print("Search keyword:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	search := strings.TrimSpace(input)

	rstfound := false
	for _, todo := range t.Todos {

		if strings.Contains(todo.title, search) {
			rstfound = true
			status := "TODO"
			if todo.status {
				status = "DONE"
			}

			fmt.Printf("[%s] #%d %-30s%s\n", status, todo.id, todo.title, todo.created.Format("02-01-2006 3:04PM"))
		}
	}

	if !rstfound {
		fmt.Println("(none)")
	}

}

func (t *TodoTui) initalMenu() {
tuiloop:
	for {
		var selectedOption string
		fmt.Println("======= TO-DO CLI =======")
		fmt.Println("1> Add task")
		fmt.Println("2> List tasks")
		fmt.Println("3> Mark task as done")
		fmt.Println("4> Delete task")
		fmt.Println("5> Search task")
		fmt.Println("6>Exit")
		fmt.Print("Choose:")
		fmt.Scanln(&selectedOption)

		switch selectedOption {
		case "1":
			t.addTask()
		case "2":
			t.listTasks()
		case "3":
			t.markTaskDone()
		case "4":
			t.deleteTask()
		case "5":
			t.searchTask()
		case "6":
			fmt.Println("Bye ..")
			break tuiloop
		default:
			fmt.Println("Invalid choice. Please select 1-6")
		}

	}

}

func (t *TodoTui) Start() {
	t.initalMenu()
}
