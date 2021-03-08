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
		UserID:       input.UserID,
		BattletagID:  input.BattletagID,
		SessonID:     input.SessonID,
		Location:     input.Location,
		Role:         input.Role,
		SrIn:         input.SrIn,
		SrOut:        input.SrOut,
		MatchOutcome: input.MatchOutcome,
	}

	qstr := `INSERT INTO game (
		userId,
		battletagId,
		sessionId,
		location,
		role,
		sr_in,
		sr_out,
		match_outcome
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	statement, err := db.Prepare(qstr)

	if err != nil {
		panic(err.Error())
	}

	res, err := statement.Exec(
		gameInput.UserID,
		gameInput.BattletagID,
		gameInput.SessonID,
		gameInput.Location,
		gameInput.Role,
		gameInput.SrIn,
		gameInput.SrOut,
		gameInput.MatchOutcome,
	)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully inserted game to session")

	lastInsertedID, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	lastInserted := db.QueryRow(`Select * from game where id=?;`, lastInsertedID)

	var (
		id, userId, battletagId, sessionId, srIn, srOut int
		location                                        model.Location
		role                                            model.Role
		matchOutcome                                    model.MatchOutcome
	)

	err = lastInserted.Scan(&id, &userId, &sessionId)

	if err != nil {
		panic(err.Error())
	}

	insertedGame := model.Game{
		ID:           id,
		UserID:       userId,
		BattletagID:  battletagId,
		SessonID:     sessionId,
		Location:     location,
		Role:         role,
		SrIn:         srIn,
		SrOut:        srOut,
		MatchOutcome: matchOutcome,
	}

	return insertedGame
}
