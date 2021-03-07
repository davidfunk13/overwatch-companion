package database

import (
	"github.com/davidfunk13/overwatch-companion/graph/model"
)

// SelectAllSessions Selects and returns all sessions from the database. This function will be refined to only select one battletag's sessions.
func SelectAllSessions() ([]*model.Session, error) {
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var data []*model.Session

	res, err := db.Query("SELECT * FROM session")

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	for res.Next() {
		var (
			id, userId int
			roleType   model.Role
		)

		err = res.Scan(&id, &userId, &roleType)

		s := model.Session{ID: id, UserID: userId, RoleType: roleType}

		data = append(data, &s)
	}

	return data, nil
}
