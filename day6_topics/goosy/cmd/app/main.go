package main

import (
	"context"
	"fmt"
	"goosy/internal/db"
	"goosy/internal/user"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/appdb?sslmode=disable"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	d, err := db.New(ctx, dbURL)

	if err != nil {
		log.Fatal(err)
	}

	us := user.NewStore(d.Pool)

	// createCtx, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel2()
	//
	// newId, err := us.Create(createCtx, "goosy1", "goosy1@mail.com")

	// newId, err := us.CreateUserWithAudit(createCtx, "goosy_tx", "goosy_tx@mail.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println("Inserted user id:", newId)

	var wg sync.WaitGroup
	for x := range 10 {
		wg.Go(func() {
			createCtx, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel2()

			_, err := us.CreateUserWithAudit(createCtx, fmt.Sprintf("user_%d", x), fmt.Sprintf("user_%d@mail.com", x))
			if err != nil {
				fmt.Println("create user with audit:", err)
			}
		})
	}

	wg.Wait()

	listCtx, cancel3 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel3()

	users, err := us.List(listCtx, 30)

	if err != nil {
		log.Fatal(err)
	}

	for _, u := range users {

		fmt.Printf("%d | %s | %s | %s\n", u.ID, u.Name, u.Email, u.CreatedAt.Format(time.RFC3339))

	}

}
