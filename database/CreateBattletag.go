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

	battletagInput := &model.Battletag{
		UserID:     input.UserID,
		Battletag:  input.Battletag,
		Platform:   &input.Platform,
		Identifier: input.Identifier,
	}

	statement, err := db.Prepare(`INSERT INTO battletag (userId, battletag, platform, identifier) VALUES (?, ?, ?, ?);`)

	if err != nil {
		panic(err.Error())
	}

	res, err := statement.Exec(battletagInput.UserID, battletagInput.Battletag, battletagInput.Platform, battletagInput.Identifier)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully inserted")

	lastInsertedID, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	lastInserted := db.QueryRow(`Select * from battletag where id=?;`, lastInsertedID)

	var userId, id int
	var battletag string
	var identifier *int
	var platform *model.Platform

	err = lastInserted.Scan(&id, &userId, &battletag, &identifier, &platform)

	insertedBattletag := model.Battletag{ID: id, UserID: userId, Battletag: battletag, Identifier: identifier, Platform: platform}

	return &insertedBattletag
}
