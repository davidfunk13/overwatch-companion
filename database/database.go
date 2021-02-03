package database

import (
	"database/sql"
	"fmt"
)

//Db holds an exported global reference to the database
var Db *sql.DB

// InitConnection esatablishes a connection to the database
func InitConnection() {
	db, err := sql.Open("mysql", "root:root@/overwatch_companion")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to database")
	Db = db
}
