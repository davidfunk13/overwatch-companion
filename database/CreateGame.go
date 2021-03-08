package database

import (
	"fmt"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

//CreateGame creates a game for the user to store against a session.
func CreateGame(input model.InputGame) model.Game {
	//open connection to the database
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	//wait until function finishes running, then close connection.
	defer db.Close()


	// NEED ATTENTION HERE 
	
	/*
		You need to get the previous game's SR here for your SR IN record. 
		Do we even need this? Why or why not?
		WHY:
			we can judge SR difference (too little, too much)
			we can more easily judge and track average sr Gained and lost.
			its better to have it than not i guess?
		WHY NOT: 
			It's more work.
			More work takes more time.
			its one more lookup/update operation before returning the result, makes it more of an expensive operation	

		if we don't lets remove srIn from the schema.
		Either way, it needs to come off of the input.  
	*/


	//END ATTENTION HERE

	// take our function input and marshal it to an InputGame Struct
	gameInput := model.InputGame{
		UserID:       input.UserID,
		BattletagID:  input.BattletagID,
		SessonID:     input.SessonID,
		Location:     input.Location,
		Role:         input.Role,
		SrOut:        input.SrOut,
		MatchOutcome: input.MatchOutcome,
	}

	//prepare to insert this input type into the database as a game
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

	//execute the insert game statement
	res, err := statement.Exec(
		gameInput.UserID,
		gameInput.BattletagID,
		gameInput.SessonID,
		gameInput.Location,
		gameInput.Role,
		gameInput.SrOut, //THIS IS TEMPORARILY REPLACING SR_IN WHILE WE DECIDE WHETHER WE LOOK UP PREVIOUS GAME SR TO UPDATE.
		gameInput.SrOut,
		gameInput.MatchOutcome,
	)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully inserted game to session")
	
	//get the id of the game record we just inserted.
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

	//we can either update the session's SR here, or after getting the id of the last inserted. be careful of where you are re-using the res, err vars

	return insertedGame
}
