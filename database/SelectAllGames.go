package database

import (
	"github.com/davidfunk13/overwatch-companion/graph/model"
)

// SelectAllGames Selects and returns all games from the database. This function will be refined to only select one sessions games.
func SelectAllGames() ([]*model.Game, error) {
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var data []*model.Game

	res, err := db.Query("SELECT * FROM game")

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	for res.Next() {
		var id, userId, sessionId int

		err = res.Scan(&id, &userId, &sessionId)

		g := model.Game{ID: id, UserID: userId, SessonID: sessionId}

		data = append(data, &g)
	}

	return data, nil
}
