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

	res, err := db.Query("SELECT * FROM battletag where ")

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	for res.Next() {
		var (
			id, userId, blizzId, level, playerLevel int
			isPublic bool
			name, urlName, portrait  string 
			platform   model.Platform
		)

		err = res.Scan(
			&id,
			&userId,
			&name,
			&urlName,
			&blizzId,
			&level,
			&playerLevel,
			&isPublic,
			&platform,
			&portrait,
		)

		b := model.Battletag{
			ID: id,
			UserID: userId,
			Name: name,
			URLName: urlName,
			BlizzID: blizzId,
			Level: level,
			PlayerLevel: playerLevel,
			IsPublic: &isPublic,
			Platform: platform,
			Portrait: portrait,
		}

		data = append(data, &b)
	}

	return data, nil
}
