package database

import (
	"database/sql"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

// GetOneBattletag returns a single battletag by userId and battletagId.
func GetOneBattletag(input *model.InputGetOneBattletag) model.QueryItemPayload {
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var (
		userId, name, urlName, portrait          string
		battletagId, blizzId, level, playerLevel int
		platform                                 model.Platform
		isPublic                                 *bool
	)

	qstr := `SELECT * FROM battletag WHERE id=? AND userId=?;`

	row := db.QueryRow(qstr, input.BattletagID, input.UserID)

	switch err := row.Scan(
		&battletagId,
		&userId,
		&name,
		&urlName,
		&blizzId,
		&level,
		&playerLevel,
		&platform,
		&isPublic,
		&portrait,
	); err {
	case sql.ErrNoRows:
		err := "No battletag found with those ids."

		payload := model.QueryItemFailure{
			Success: false,
			Error:   &err,
		}
		return payload
	case nil:
		battletag := model.Battletag{
			ID:          battletagId,
			UserID:      userId,
			Name:        name,
			URLName:     urlName,
			BlizzID:     blizzId,
			Level:       level,
			PlayerLevel: playerLevel,
			Platform:    platform,
			IsPublic:    isPublic,
			Portrait:    portrait,
		}

		payload := model.QueryItemSuccess{
			Success: true,
			Data:    battletag,
		}

		return payload
	default:
		panic(err)
	}
}
