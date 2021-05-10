package database

import (
	"github.com/davidfunk13/overwatch-companion/graph/model"
)

// SelectAllBattletags Selects and returns all battletags from the database. This function will be refined to only fetch a user's battletags.
func GetAllBattletags(uId string) []*model.Battletag {
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var data []*model.Battletag

	res, err := db.Query("SELECT * FROM battletag where userId=?", uId)

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	for res.Next() {
		var (
			userId, name, urlName, portrait, created_at, updated_at string
			id, blizzId, level, playerLevel                         int
			isPublic                                                bool
			platform                                                model.Platform
		)

		err = res.Scan(
			&id,
			&userId,
			&name,
			&urlName,
			&blizzId,
			&level,
			&playerLevel,
			&platform,
			&isPublic,
			&portrait,
			&created_at,
			&updated_at,
		)

		if err != nil {
			panic(err.Error())
		}

		b := model.Battletag{
			ID:          id,
			UserID:      userId,
			Name:        name,
			URLName:     urlName,
			BlizzID:     blizzId,
			Level:       level,
			PlayerLevel: playerLevel,
			Platform:    platform,
			IsPublic:    &isPublic,
			Portrait:    portrait,
			CreatedAt:   created_at,
			UpdatedAt:   &updated_at,
		}

		data = append(data, &b)
	}

	return data
}
