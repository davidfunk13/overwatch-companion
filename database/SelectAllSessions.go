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
		var (
			id, userId, battletagId int
		)

		err = res.Scan(&id, &userId, &battletagId)

		s := model.Session{ID: id, UserID: userId, BattletagID: battletagId}

		data = append(data, &s)
	}

	return data, nil
}
