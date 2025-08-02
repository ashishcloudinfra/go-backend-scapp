package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
	"server.simplifycontrol.com/secrets"
)

var DB *sql.DB

// InitializeDatabase sets up the database connection
func InitializeDatabase() {
	var err error

	// load from ENV VARIABLE
	connStr := secrets.SecretJSON.Database.URL
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return
	}

	// Ping the database to verify the connection
	if err = DB.Ping(); err != nil {
		log.Printf("Error pinging database: %v", err)
		return
	}
	log.Println("Database connection established")
}

// CloseDatabase closes the database connection
func CloseDatabase() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}
}

// reconnectToDatabase attempts to reconnect to the database
func reconnectToDatabase() bool {
	var err error
	// Retry logic to reconnect
	for {
		err = DB.Ping() // Attempt to ping the existing DB connection
		if err == nil {
			log.Println("Successfully reconnected to the database.")
			return true
		}

		log.Printf("Error reconnecting to database: %v", err)
		time.Sleep(5 * time.Second) // Retry after a delay
	}
}
