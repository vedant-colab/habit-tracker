package database

import (
	"database/sql"
	"fmt"
	"habit-tracker/config"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	dsn := config.GetEnv("DATABASE_DSN", "")
	// dsn := os.Getenv("DATABASE_DSN")

	// Open a connection
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Verify the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Successfully connected to the database")
}
