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

	// take our function input and marshal it to an InputGame Struct
	gameInput := model.InputGame{
		UserID:       input.UserID,
		BattletagID:  input.BattletagID,
		SessionID:    input.SessionID,
		Location:     input.Location,
		Role:         input.Role,
		SrOut:        input.SrOut,
		MatchOutcome: input.MatchOutcome,
	}

	// define set of variables to hold previous game's values
	var (
		prevId, prevUserId, prevBattletagId, prevSessionId, prevSrIn, prevSrOut int
		prevLocation                                        model.Location
		prevRole                                            model.Role
		prevMatchOutcome                                    model.MatchOutcome
	)
	
	//Get games for this role in this user's battletags' session, if they exist.
	prevQstr := `SELECT * from game WHERE userId=? AND battletagId=? AND sessionId=? AND role=?;`
	prevGameStatement, err := db.Query(prevQstr, input.UserID, input.BattletagID, input.SessionID, input.Role)

	if err != nil {
		panic(err.Error())
	}

	//create empty array to hold temporary existing games, if they exist. 
	var prevGames []*model.Game

	for prevGameStatement.Next() {
		err = prevGameStatement.Scan(&prevId, &prevUserId, &prevBattletagId, &prevSessionId, &prevLocation, &prevRole, &prevSrIn, &prevSrOut, &prevMatchOutcome)

		g := model.Game{
			ID: prevId, 
			UserID: prevUserId, 
			BattletagID: prevBattletagId, 
			SessionID: prevSessionId,
			Location: prevLocation,
			Role: prevRole,
			SrIn: prevSrIn,
			SrOut: prevSrOut,
			MatchOutcome: prevMatchOutcome,
		}

		prevGames = append(prevGames, &g)		
	}
	

	// if we get any games back from the db (meaning there is a previous game to pull srIn from...)
	var prevGameSrOut int
	
	if len(prevGames) > 0 {
		lastGame:= prevGames[len(prevGames) - 1]
		prevGameSrOut = lastGame.SrOut
		fmt.Println("Last Game: ", lastGame)
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
		gameInput.SessionID,
		gameInput.Location,
		gameInput.Role,
		prevGameSrOut, //THIS IS TEMPORARILY REPLACING SR_IN WHILE WE DECIDE WHETHER WE LOOK UP PREVIOUS GAME SR TO UPDATE.
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

	// get the game we just inserted
	lastInserted := db.QueryRow(`Select * from game where id=?;`, lastInsertedID)

	// define set of variables to hold values of game we just inserted./ 
	var (
		id, userId, battletagId, sessionId, srIn, srOut int 
		role model.Role
		location model.Location
		matchOutcome model.MatchOutcome
	)

	err = lastInserted.Scan(&id, &userId, &battletagId, &sessionId, &location, &role, &srIn, &srOut, &matchOutcome )

	if err != nil {
		panic(err.Error())
	}

	// marshal variables to game struct representing inserted game record.
	insertedGame := model.Game{
		ID:           id,
		UserID:       userId,
		BattletagID:  battletagId,
		SessionID:    sessionId,
		Location:     location,
		Role:         role,
		SrIn:         srIn,
		SrOut:        srOut,
		MatchOutcome: matchOutcome,
	}
	
	var updateSessionStr string
	
	switch role {
	case "TANK":
		updateSessionStr = "UPDATE session SET sr_tank=? WHERE id=? and userId=? and battletagId=?;"
	case "DAMAGE":
		updateSessionStr = "UPDATE session SET sr_damage=? WHERE id=? and userId=? and battletagId=?;"
	case "SUPPORT":
		updateSessionStr = "UPDATE session SET sr_support=? WHERE id=? and userId=? and battletagId=?;"
	}

	//try passing in a pointer to the string here, observe behavior. 
	updateSessionStatement, err := db.Prepare(updateSessionStr)

	// prevGameSrOut will be our value we use to update, and will be presumably going to the right place because of the above switch statement.
	// srOut will also work here, as it is the same value
	updateSessionStatement.Exec(
		srOut,
		sessionId,
		userId,
		battletagId,		
	)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Session SR for %s sucessfully updated in session %d for battletag %d for user %d", role, sessionId, battletagId, userId)

	return insertedGame
}
