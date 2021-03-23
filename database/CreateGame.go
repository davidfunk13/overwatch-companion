package database

import (
	"fmt"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

//CreateGame creates a game for the user to store against a session.
func CreateGame(input model.InputGame) model.MutateItemPayload {

	//payload
	var payload model.MutateItemPayload

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
		prevUserId                                                  string
		prevId, prevBattletagId, prevSessionId, prevSrIn, prevSrOut int
		prevLocation                                                model.Location
		prevRole                                                    model.Role
		prevMatchOutcome                                            model.MatchOutcome
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
			ID:           prevId,
			UserID:       prevUserId,
			BattletagID:  prevBattletagId,
			SessionID:    prevSessionId,
			Location:     prevLocation,
			Role:         prevRole,
			SrIn:         prevSrIn,
			SrOut:        prevSrOut,
			MatchOutcome: prevMatchOutcome,
		}

		prevGames = append(prevGames, &g)
	}

	// if we get any games back from the db (meaning there is a previous game to pull srIn from...)
	// grab the sr_out of the most recent game.
	var prevGameSrOut int

	if len(prevGames) > 0 {
		lastGame := prevGames[len(prevGames)-1]
		prevGameSrOut = lastGame.SrOut
		fmt.Println("Last Game: ", lastGame)
		// ..here we check if the starting SR is > 0. if it's zero, we bail.
		// ()
	}

	//prepare to insert input type into the database as a new game
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

	//execute the insert game statement, with the previous game's sr_out value as the new game's sr_in value
	res, err := statement.Exec(
		gameInput.UserID,
		gameInput.BattletagID,
		gameInput.SessionID,
		gameInput.Location,
		gameInput.Role,
		prevGameSrOut,
		gameInput.SrOut,
		gameInput.MatchOutcome,
	)

	if err != nil {
		errorString := err.Error()

		payload = model.MutateItemPayloadFailure{
			Success: false,
			Error:   "Error inserting game into database",
			Data:    &errorString,
		}

		return payload
	}

	fmt.Println("Successfully inserted game to session")

	//get the id of the new game record we just inserted, store for reference.
	lastInsertedID, err := res.LastInsertId()

	if err != nil {
		errorString := err.Error()

		payload = model.MutateItemPayloadFailure{
			Success: false,
			Error:   "Error getting Id of last inserted game",
			Data:    &errorString,
		}

		return payload
	}

	// get the game we just inserted
	lastInserted := db.QueryRow(`Select * from game where id=?;`, lastInsertedID)

	// define new set of variables to hold values of game we just inserted.
	var (
		userId                                  string
		id, battletagId, sessionId, srIn, srOut int
		role                                    model.Role
		location                                model.Location
		matchOutcome                            model.MatchOutcome
	)

	err = lastInserted.Scan(&id, &userId, &battletagId, &sessionId, &location, &role, &srIn, &srOut, &matchOutcome)

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

	//variable to hold our statement to update the session's SR, depending on what role the game we just created is.
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

	// srOut will be our value we use to update, and will be presumably going to the right place because of the above switch statement.
	updateSessionStatement.Exec(
		srOut,
		sessionId,
		userId,
		battletagId,
	)

	if err != nil {
		errorString := err.Error()

		payload = model.MutateItemPayloadFailure{
			Success: false,
			Error:   "Error updating session SR after this game inserted.",
			Data:    &errorString,
		}

		// We need to delete the game we just added and fail the operation as a whole because we need
		// to keep our SR values consistnent with what is in the game tables.
		lastGameID := int(lastInsertedID)

		DeleteGame(&lastGameID)

		return payload
	}

	fmt.Printf("Session SR for %s sucessfully updated in session %d for battletag %d for user %d", role, sessionId, battletagId, userId)

	payload = &model.MutateItemPayloadSuccess{
		ID:      insertedGame.ID,
		Success: true,
		Message: "Successfully inserted game.",
		Data:    insertedGame,
	}

	return payload
}
