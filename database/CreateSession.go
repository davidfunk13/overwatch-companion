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
		BattletagID: input.BattletagID,
	}

	statement, err := db.Prepare(`INSERT INTO session (userId, battletagId) VALUES (?, ?);`)

	if err != nil {
		panic(err.Error())
	}

	res, err := statement.Exec(sessionInput.UserID, sessionInput.BattletagID)

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
		id, userId, battletagId int
	)
	err = lastInserted.Scan(&id, &userId, &battletagId)

	insertedSession := model.Session{
		ID:       id,
		UserID:   id,
		BattletagID: battletagId,
	}

	return insertedSession
}
