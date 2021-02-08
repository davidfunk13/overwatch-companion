package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

//Db holds an exported global reference to the database
var Db *sql.DB

// InitConnection esatablishes a connection to the database
func InitConnection() {
	env := os.Getenv("APP_ENV")
	var connStr string
	if env == "production" {
		fmt.Println("Connected to database")
		connStr = os.Getenv("JAWSDB_URL")
	} else {
		connStr = "root:@/overwatch_companion"
	}

	fmt.Printf("{ env: %s, connectionString: %s }\n", env, connStr)

	db, err := sql.Open("mysql", connStr)

	if err != nil {
		panic(err.Error())
	}

	Db = db
}

// SelectAllUsers Selects and returns all users from the database. This is a function for development purposes only.
func SelectAllUsers() []*model.User {

	// declare empty slice representing array of users
	var data []*model.User

	// run Query for 2+ rows, QueryRow for 1.
	res, err := Db.Query("SELECT * FROM user")

	// keep looping as long as res.Next() returns true
	for res.Next() {
		// declare variables for the incoming cols values
		var id, name, email string

		// provide pointers to col value variables to .Scan, will parse and assign col data to variable.
		err = res.Scan(&id, &name, &email)

		//assign values to user model
		user := model.User{ID: id, Name: name, Email: email}

		//send this new user to the data array
		data = append(data, &user)
	}

	// if there is an error making the query, panic.
	if err != nil {
		fmt.Println("hit this err")
		panic(err.Error())
	}

	return data
}

// func SelectOneUser() {
// 	// Selecting one row
// 	/*
// 		id := 1
// 		var col string
// 		sqlStatement := `SELECT col FROM my_table WHERE id=$1`
// 		row := db.QueryRow(sqlStatement, id)
// 		err := row.Scan(&col)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				fmt.Println("Zero rows found")
// 			} else {
// 				panic(err)
// 			}
// 		}
// 	*/

// }
