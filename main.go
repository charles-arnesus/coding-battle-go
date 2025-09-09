package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to Postgres
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to Postgres:", err)
	}

	fmt.Println("Connected to Postgres 🎉")

	// Interactive loop
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your name (or 'exit' to quit): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		name := strings.TrimSpace(input)

		if name == "exit" {
			fmt.Println("Bye! 👋")
			break
		}

		fmt.Printf("Hello, %s! 👋\n", name)

		// Save to Postgres
		_, err = db.Exec("INSERT INTO users(name) VALUES($1)", name)
		if err != nil {
			fmt.Println("Failed to save to DB:", err)
		} else {
			fmt.Println("Saved to database ✅")
		}
	}
}
