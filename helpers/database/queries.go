package database

import (
	"database/sql"
	"errors"
	"log"
)

// Query executes a SQL query and returns rows
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	// Try executing the query
	rows, err := DB.Query(query, args...)
	if err != nil {
		// If the error is a connection issue, try reconnecting
		if isConnectionError(err) {
			log.Println("Connection lost, attempting to reconnect...")
			if reconnectToDatabase() {
				// Retry the query after successful reconnection
				rows, err = DB.Query(query, args...)
			}
		}
		// If still an error, return it
		if err != nil {
			log.Printf("Error executing query after reconnect: %v", err)
			return nil, err
		}
	}
	return rows, nil
}

// Execute runs a non-select query (INSERT, UPDATE, DELETE)
func Execute(query string, args ...interface{}) (sql.Result, error) {
	// Try executing the query
	result, err := DB.Exec(query, args...)
	if err != nil {
		// If the error is a connection issue, try reconnecting
		if isConnectionError(err) {
			log.Println("Connection lost, attempting to reconnect...")
			if reconnectToDatabase() {
				// Retry the query after successful reconnection
				result, err = DB.Exec(query, args...)
			}
		}
		// If still an error, return it
		if err != nil {
			log.Printf("Error executing query after reconnect: %v", err)
			return nil, err
		}
	}
	return result, nil
}

// isConnectionError checks if the error is related to a database connection issue
func isConnectionError(err error) bool {
	// In Go's SQL package, certain error types indicate connection issues.
	// Here, we check for specific errors that suggest the connection was lost.
	return errors.Is(err, sql.ErrConnDone) || err.Error() == "driver: bad connection"
}
