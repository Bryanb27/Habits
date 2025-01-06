package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDatabaseExample() *sql.DB {
	var (
		username = "yourUsername"
		password = "yourPassword"
		dbname   = "yourDBName"
	)

	// Build the connection string
	connStr := fmt.Sprintf("user=%s password=%s host=localhost port=5432 dbname=%s sslmode=disable", username, password, dbname)

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}

	// Check if the connection works
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")

	return db
}
