package database

import (
	"fmt"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

//CreateGame creates a game for the user to store against a session.
func CreateGame(input model.InputGame) model.Game {
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	gameInput := model.InputGame{
		UserID: input.UserID,
		SessionID: input.SessionID,
	}

	statement, err := db.Prepare(`INSERT INTO game (userId, sessionId) VALUES (?, ?);`)

	if err != nil {
		panic(err.Error())
	}

	res, err := statement.Exec(gameInput.UserID, gameInput.SessionID)
	
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully inserted game to session")

	lastInsertedID, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}
	
	lastInserted := db.QueryRow(`Select * from game where id=?;`, lastInsertedID)

	var userId, id, sessionId int

	err = lastInserted.Scan(&id, &userId, &sessionId)
	
	if err != nil {
		panic(err.Error())
	}

	insertedGame := model.Game{
		ID: id,
		UserID: userId,
		SessonID: sessionId,
	}

	return insertedGame
}