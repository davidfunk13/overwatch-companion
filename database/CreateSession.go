package database

import (
	"fmt"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

//CreateSession creates a game session for the user to store and track games and stats against. Holds a particular role.
func CreateSession(input model.InputSession) model.Session {
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	sessionInput := model.InputSession{
		UserID:   input.UserID,
		RoleType: input.RoleType,
	}

	statement, err := db.Prepare(`INSERT INTO session (userId, roleType) VALUES (?, ?);`)

	if err != nil {
		panic(err.Error())
	}

	res, err := statement.Exec(sessionInput.UserID, sessionInput.RoleType)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully inserted Session record")

	lastInsertedID, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	lastInserted := db.QueryRow(`Select * from session where id=?;`, lastInsertedID)

	var (
		userId, id int
		roleType   model.Role
	)
	err = lastInserted.Scan(&id, &userId, &roleType)

	insertedSession := model.Session{
		ID:       id,
		UserID:   id,
		RoleType: roleType,
	}

	return insertedSession
}
