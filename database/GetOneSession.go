package database

import (
	"database/sql"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

// Get session gets a single session by id and battletagId
func GetSession(input *model.InputGetOneSessionByIDAndBattletagID) model.QueryPayload {
	//open connection to the database
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	//wait until function finishes running, then close connection.
	defer db.Close()

	var (
		userId                                                                                                     string
		id, battletagId, starting_sr_tank, sr_tank, starting_sr_damage, sr_damage, starting_sr_support, sr_support int
	)

	qstr := `SELECT * FROM session WHERE id=? AND battletagId=?;`

	row := db.QueryRow(qstr, input.ID, input.BattletagID)

	switch err := row.Scan(
		&id,
		&userId,
		&battletagId,
		&starting_sr_tank,
		&sr_tank,
		&starting_sr_damage,
		&sr_damage,
		&starting_sr_support,
		&sr_support,
	); err {
	case sql.ErrNoRows:
		payload := model.GetOneItemPayloadFailure{
			Success: false,
			Error:   "No session found with those ids.",
		}

		return payload
	case nil:

		session := model.Session{
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

		payload := model.GetOneItemPayloadSuccess{
			Success: true,
			Data:    session,
		}

		return payload
	default:
		panic(err)
	}
}