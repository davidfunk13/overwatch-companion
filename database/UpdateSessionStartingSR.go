package database

import (
	"github.com/davidfunk13/overwatch-companion/graph/model"
)

//THIS ROUTE SHOULD ONLY WORK IF THE SESSION DOES NOT CONTAIN ANY GAMES!

// UpdateSessionStartingSR takes a role as inupt and a new starting SR value along with a session id.
func UpdateSessionStartingSR(input model.InputUpdateSessionStartingSr) model.MutateItemPayload {
	// variable to store our payload
	var payload model.MutateItemPayload
	//open connection to the database
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	//wait until function finishes running, then close connection.
	defer db.Close()

	getHasGames := &model.InputGetGames{
		UserID:      input.UserID,
		BattletagID: input.BattletagID,
		SessionID:   input.ID,
		Role:        &input.Role,
	}

	hasGames, err := SelectAllGames(*getHasGames)

	if err != nil {
		errorString := err.Error()

		payload = model.MutateItemPayloadFailure{
			Success: false,
			Error:   "Error checking if this session has any games for the requested role.",
			Data:    &errorString,
		}

		return payload
	}

	if len(hasGames) > 0 {
		errorString := "This session already has games of this role type in it. You cannot update the starting sr of a role that has a game in it."

		payload = model.MutateItemPayloadFailure{
			Success: false,
			Error:   "Error checking if this session has any games for the requested role.",
			Data:    &errorString,
		}

		return payload
	}

	var updateSessionStr string

	switch input.Role {
	case "TANK":
		updateSessionStr = "UPDATE session SET sr_tank=? WHERE id=? and userId=? and battletagId=?;"
	case "DAMAGE":
		updateSessionStr = "UPDATE session SET sr_damage=? WHERE id=? and userId=? and battletagId=?;"
	case "SUPPORT":
		updateSessionStr = "UPDATE session SET sr_support=? WHERE id=? and userId=? and battletagId=?;"
	default:
		payload = &model.MutateItemPayloadFailure{
			Success: false,
			Error:   "Failure to get Role from input in UpdateSessionStartingSR",
		}

		return payload
	}

	//try passing in a pointer to the string here, observe behavior.
	updateSessionStatement, err := db.Prepare(updateSessionStr)

	// srOut will be our value we use to update, and will be presumably going to the right place because of the above switch statement.
	res, err := updateSessionStatement.Exec(
		input.StartingSr,
		input.ID,
		input.UserID,
		input.BattletagID,
	)

	if err != nil {
		errorString := err.Error()

		payload = model.MutateItemPayloadFailure{
			Success: false,
			Error:   "Error updating session SR.",
			Data:    &errorString,
		}

		return payload
	}

	//get the id of the new session record we just updated, store for reference.
	lastInsertedID, err := res.LastInsertId()

	if err != nil {
		errorString := err.Error()

		payload = model.MutateItemPayloadFailure{
			Success: false,
			Error:   "Error getting ID of session just updated.",
			Data:    &errorString,
		}

		return payload
	}

	// get the session we just updated
	lastInserted := db.QueryRow(`Select * from session where id=?;`, lastInsertedID)

	// define new set of variables to hold values of game we just inserted.
	var id, userId, battletagId, starting_sr_tank, starting_sr_damage, starting_sr_support, sr_tank, sr_damage, sr_support int

	err = lastInserted.Scan(
		&id,
		&userId,
		&battletagId,
		&starting_sr_tank,
		&starting_sr_damage,
		&starting_sr_support,
		&sr_tank,
		&sr_damage,
		&sr_support,
	)

	if err != nil {
		panic(err.Error())
	}

	// marshal variables to game struct representing inserted game record.
	updatedSesh := model.Session{
		ID:                id,
		UserID:            userId,
		BattletagID:       battletagId,
		StartingSrTank:    starting_sr_tank,
		SrTank:            sr_tank,
		StartingSrDamage:  starting_sr_damage,
		SrDamage:          sr_damage,
		StartingSrSupport: starting_sr_support,
		SrSupport:         sr_support,
	}

	payload = &model.MutateItemPayloadSuccess{
		ID:      id,
		Success: true,
		Message: string("Sucessfully updated session starting sr for " + input.Role),
		Data:    updatedSesh,
	}

	return payload
}
