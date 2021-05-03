package database

import (
	"database/sql"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

// SelectAllGames Selects and returns all games from the database. This function will be refined to only select one sessions games.
func GetAllGames(input model.InputGetGames) model.QueryPayload {
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var data []*model.Game

	var (
		qstr string
		res  *sql.Rows
		qErr error
	)

	if input.Role == nil {
		qstr = "SELECT * FROM game WHERE userId=? AND battletagId=? AND sessionId=?"
		res, qErr = db.Query(qstr, input.UserID, input.BattletagID, input.SessionID)
	} else {
		qstr = "SELECT * FROM game WHERE userId=? AND battletagId=? AND sessionId=? AND role=?"
		res, qErr = db.Query(qstr, input.UserID, input.BattletagID, input.SessionID, input.Role)
	}

	if qErr != nil {

		payload := model.GetAllGamesPayloadFailure{
			Success: false,
			Error:   "Error getting games list for this user's battletag's session.",
		}

		return payload
	}

	defer res.Close()

	for res.Next() {

		var (
			userId                                    string
			id, battletagId, sessionId, sr_in, sr_out int
			location                                  model.Location
			role                                      model.Role
			matchOutcome                              model.MatchOutcome
		)

		err = res.Scan(&id, &userId, &battletagId, &sessionId, &location, &role, &sr_in, &sr_out, &matchOutcome)

		if err != nil {
			payload := model.GetAllGamesPayloadFailure{
				Success: false,
				Error:   "Error getting games list for this user's battletag's session.",
			}

			return payload
		}

		g := model.Game{
			ID:           id,
			UserID:       userId,
			BattletagID:  battletagId,
			SessionID:    sessionId,
			Location:     location,
			Role:         role,
			SrIn:         sr_in,
			SrOut:        sr_out,
			MatchOutcome: matchOutcome,
		}

		data = append(data, &g)
	}
	payload := model.GetAllGamesPayloadSuccess{
		Success: true,
		Message: "All games for this user's battletag's session have been successfully fetched.",
		Data:    data,
	}

	return payload
}
