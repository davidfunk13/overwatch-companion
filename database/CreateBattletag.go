package database

import (
	"fmt"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

// CreateBattletag is a function that creates a new battletag and returns it to the user.
func CreateBattletag(input model.InputBattletag) *model.Battletag {
	db, err := OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	battletagInput := &model.InputBattletag{
		UserID:     input.UserID,
		Name:  input.Name,
		URLName:   input.URLName,
		BlizzID: input.BlizzID,
		Level: input.Level,
		PlayerLevel: input.PlayerLevel,
		IsPublic: input.IsPublic,
		Platform: input.Platform,
		Portrait: input.Portrait,
	}

	statement, err := db.Prepare(`INSERT INTO battletag (userId, name, urlName, blizzId, level, playerLevel, platform, isPublic, portrait) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`)

	if err != nil {
		panic(err.Error())
	}

	res, err := statement.Exec(
		battletagInput.UserID,
		battletagInput.Name,
		battletagInput.URLName,
		battletagInput.BlizzID,
		battletagInput.Level,
		battletagInput.PlayerLevel,
		battletagInput.IsPublic,
		battletagInput.Platform,
		battletagInput.Portrait,
	)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully inserted")

	lastInsertedID, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	lastInserted := db.QueryRow(`Select * from battletag where id=?;`, lastInsertedID)

	var (
		userId, id, blizzId, level, playerLevel int
		isPublic bool
		name, urlName, portrait  string 
		platform   model.Platform
	)
	err = lastInserted.Scan(&id, &userId, &name, &urlName, &blizzId, &level, &playerLevel, &isPublic, &platform, &portrait )

	insertedBattletag := model.Battletag{
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

	return &insertedBattletag
}