package database

import (
	"database/sql"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

// Get session gets a single session by id and battletagId
func GetCurrentRole(input *model.InputGetCurrentRole) *model.CurrentRole {
	//open connection to the database
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	//wait until function finishes running, then close connection.
	defer db.Close()

	var (
		userId, created_at, updated_at, qstr                                                                             string
		id, battletagId, starting_sr_tank, sr_tank, starting_sr_damage, sr_damage, starting_sr_support, sr_support int
		row *sql.Row
	)

	qstr = `SELECT * FROM session WHERE id=? LIMIT 1;`
	row = db.QueryRow(qstr, input.SessionID)
	
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
		return &model.CurrentRole{}
	case nil:
		switch input.Role {
				case "TANK":
				role := model.CurrentRole{
					Role: &input.Role,
					CurrentSr: &sr_tank,
					StartingSr: &starting_sr_tank,
				}
				return &role
				
				case "DAMAGE": 
				
				role := model.CurrentRole{
					Role: &input.Role,
					CurrentSr: &sr_damage,
					StartingSr: &starting_sr_damage,
				}
				return &role
				
				case "SUPPORT": 
				role := model.CurrentRole{
					Role: &input.Role,
					CurrentSr: &sr_support,
					StartingSr: &starting_sr_support,
				}
				return &role
			default: 
		panic("No role detected")
		}
	default:
		panic(err)
	}
}
