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

	db, err := sql.Open("mysql", connStr)

	if err != nil {
		panic(err.Error())
	}

	Db = db
}

// SelectAllUsers Selects and returns all users from the database. This is a function for development purposes only.
func SelectAllUsers() []*model.User {
	var data []*model.User

	res, err := Db.Query("SELECT * FROM user")

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	for res.Next() {
		var id, name, email string

		err = res.Scan(&id, &name, &email)

		user := model.User{ID: id, Name: name, Email: email}

		data = append(data, &user)
	}

	return data
}

// SelectAllBattletags Selects and returns all battletags that belong to a user from the database.
func SelectAllBattletags() []*model.Battletag {
	var data []*model.Battletag

	res, err := Db.Query("SELECT * FROM battletag")

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	for res.Next() {
		var ID, userID int
		var identifier *int
		var battletag string
		var platform *model.Platform

		err = res.Scan(&ID, &userID, &battletag, &identifier, &platform)

		b := model.Battletag{ID: ID, Battletag: battletag, UserID: userID, Identifier: identifier, Platform: platform}

		data = append(data, &b)
	}

	return data
}
