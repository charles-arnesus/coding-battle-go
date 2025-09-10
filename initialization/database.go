package initialization

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func dbConnection() *sql.DB {
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

	return db
}
