package database

import (
	"github.com/davidfunk13/overwatch-companion/graph/model"
)

// SelectAllSessions Selects and returns all sessions from the database. This function will be refined to only select one battletag's sessions.
func SelectAllSessions(uid int, bid int) ([]*model.Session, error) {
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var data []*model.Session

	res, err := db.Query("SELECT * FROM session where userId=? AND battletagId=?", uid, bid)

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	for res.Next() {
		var id, userId, battletagId, starting_sr_tank, sr_tank, starting_sr_damage, sr_damage, starting_sr_support, sr_support int

		err = res.Scan(
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

		s := model.Session{
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

		data = append(data, &s)
	}

	return data, nil
}
