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
		UserID:            input.UserID,
		BattletagID:       input.BattletagID,
		StartingSrTank:    input.StartingSrTank,
		StartingSrDamage:  input.StartingSrDamage,
		StartingSrSupport: input.StartingSrSupport,
	}
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

	res, err := statement.Exec(
		sessionInput.UserID,
		sessionInput.BattletagID,
		sessionInput.StartingSrTank,
		sessionInput.StartingSrTank,
		sessionInput.StartingSrDamage,
		sessionInput.StartingSrDamage,
		sessionInput.StartingSrSupport,
		sessionInput.StartingSrSupport,
	)

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
		id, userId, battletagId, starting_sr_tank, sr_tank, starting_sr_damage, sr_damage, starting_sr_support, sr_support int
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
	}

	return insertedSession
}
