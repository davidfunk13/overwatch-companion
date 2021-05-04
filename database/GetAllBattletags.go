package database

import (
	"github.com/davidfunk13/overwatch-companion/graph/model"
)

// SelectAllBattletags Selects and returns all battletags from the database. This function will be refined to only fetch a user's battletags.
func GetAllBattletags(uId string) model.QueryPayload {
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
			userId                          string
			id, blizzId, level, playerLevel int
			isPublic                        bool
			name, urlName, portrait         string
			platform                        model.Platform
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
		)

		if err != nil {
			payload := model.GetAllBattletagsPayloadFailure{
				Success: false,
				Error:   "Failed getting list of battletags from the database.",
				Data:    make([]*model.Battletag, 0),
			}
			return payload
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
		}

		data = append(data, &b)
	}

	payload := model.GetAllBattletagsPayloadSuccess{
		Success: true,
		Message: "All battletags for this user have been fetched",
		Data:    data,
	}

	return payload
}
