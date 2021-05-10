package database

import (
	"fmt"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

//CreateSession creates a game session for the user to store and track games and stats against. Holds a particular role.
func CreateSession(input model.InputSession) model.MutateItemPayload {

	//payload
	var payload model.MutateItemPayload

	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//table has value of int not null, if omitted, we want the starting SR values to be int types that will zero to 0.
	var input_starting_sr_tank, input_starting_sr_damage, input_starting_sr_support int

	//if we include one starting sr and not the others, make sure that works too.
	if input.StartingSrTank != nil {
		input_starting_sr_tank = *input.StartingSrTank
	}

	if input.StartingSrDamage != nil {
		input_starting_sr_damage = *input.StartingSrDamage
	}

	if input.StartingSrSupport != nil {
		input_starting_sr_support = *input.StartingSrSupport
	}

	//marshal new altered input to InputSession struct
	sessionInput := model.InputSession{
		UserID:            input.UserID,
		BattletagID:       input.BattletagID,
		StartingSrTank:    &input_starting_sr_tank,
		StartingSrDamage:  &input_starting_sr_damage,
		StartingSrSupport: &input_starting_sr_support,
	}

	// Prepare statement to insert new session into db
	qstr := `INSERT INTO session (
		userId,
		battletagId,
		starting_sr_tank,
		sr_tank,
		starting_sr_damage,
		sr_damage, 
		starting_sr_support,
		sr_support
	 ) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	statement, err := db.Prepare(qstr)

	if err != nil {
		panic(err.Error())
	}

	// Insert new statement into database
	res, err := statement.Exec(
		sessionInput.UserID,
		sessionInput.BattletagID,
		sessionInput.StartingSrTank,
		0,
		sessionInput.StartingSrDamage,
		0,
		sessionInput.StartingSrSupport,
		0,
	)

	if err != nil {
		errorString := err.Error()

		payload = model.MutateItemPayloadFailure{
			Success: false,
			Error:   "Error creating session",
			Data:    &errorString,
		}

		return payload
	}

	fmt.Println("Successfully inserted Session record")

	// Get last inserted sessions id
	lastInsertedID, err := res.LastInsertId()

	if err != nil {
		errorString := err.Error()

		payload = model.MutateItemPayloadFailure{
			Success: false,
			Error:   "Error getting Id of last inserted session",
			Data:    &errorString,
		}

		return payload
	}

	// Get last inserted session
	lastInserted := db.QueryRow(`Select * from session where id=?;`, lastInsertedID)

	var (
		userId, created_at, updated_at                                                                             string
		id, battletagId, starting_sr_tank, sr_tank, starting_sr_damage, sr_damage, starting_sr_support, sr_support int
	)

	err = lastInserted.Scan(
		&id,
		&userId,
		&battletagId,
		&starting_sr_tank,
		&sr_tank,
		&starting_sr_damage,
		&sr_damage,
		&starting_sr_support,
		&sr_support,
		&created_at,
		&updated_at,
	)

	insertedSession := model.Session{
		ID:                id,
		UserID:            userId,
		BattletagID:       battletagId,
		StartingSrTank:    starting_sr_tank,
		SrTank:            sr_tank,
		StartingSrDamage:  starting_sr_damage,
		SrDamage:          sr_damage,
		StartingSrSupport: starting_sr_support,
		SrSupport:         sr_support,
		CreatedAt:         created_at,
		UpdatedAt:         &updated_at,
	}

	payload = model.MutateItemPayloadSuccess{
		ID:      insertedSession.ID,
		Success: true,
		Message: "Successfully inserted a new session into the database.",
		Data:    insertedSession,
	}

	return payload
}
