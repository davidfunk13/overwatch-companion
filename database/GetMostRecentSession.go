package database

import (
	"database/sql"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

// Get session gets a single session by id and battletagId
func GetMostRecentSession(input *model.InputGetMostRecentSession) []*model.Session {
	//open connection to the database
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	//wait until function finishes running, then close connection.
	defer db.Close()

	var (
		userId, created_at, updated_at                                                                             string
		id, battletagId, starting_sr_tank, sr_tank, starting_sr_damage, sr_damage, starting_sr_support, sr_support int
	)

	qstr := `SELECT * FROM session WHERE userId=? AND battletagId=? ORDER BY created_at DESC LIMIT 1;`

	row := db.QueryRow(qstr, input.UserID, input.BattletagID)

	var returnSlice []*model.Session

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
		&created_at,
		&updated_at,
	); err {
	case sql.ErrNoRows:
		return returnSlice
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
			CreatedAt:         created_at,
			UpdatedAt:         &updated_at,
		}

		returnSlice = append(returnSlice, &session)

		return returnSlice
	default:
		panic(err)
	}
}
