package database

import (
	"github.com/davidfunk13/overwatch-companion/graph/model"
)

// SelectAllBattletags Selects and returns all battletags from the database. This function will be refined to only fetch a user's battletags.
func SelectAllBattletags() ([]*model.Battletag, error) {
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var data []*model.Battletag

	res, err := db.Query("SELECT * FROM battletag")

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	for res.Next() {
		var ID, userID int
		var identifier *int
		var battletag string
		var platform model.Platform

		err = res.Scan(&ID, &userID, &battletag, &platform,  &identifier)

		b := model.Battletag{ID: ID, Battletag: battletag, UserID: userID, Identifier: identifier, Platform: platform}

		data = append(data, &b)
	}

	return data, nil
}
