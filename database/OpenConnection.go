package database

import (
	"database/sql"
	"fmt"
	"os"
)

// OpenConnection esatablishes a connection to the database
func OpenConnection() (DB *sql.DB, err error) {
	env := os.Getenv("APP_ENV")

	var connStr string

	if env == "production" {
		fmt.Println("Connected to database")
		connStr = os.Getenv("JAWSDB_URL")
	} else {
		connStr = "root:@/overwatch_companion"
	}

	db, err := sql.Open("mysql", connStr)

	if err != nil {
		fmt.Println("Error connecting to database.")
		return nil, err
	}

	// Db = db
	return db, nil
}
