package database

import (
	"strconv"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

//DeleteGame deletes a single game from the database by id
func DeleteGame(id *int) (model.MutateItemPayload, error) {
	db, err := OpenConnection()

	defer db.Close()
	statement, err := db.Prepare("DELETE FROM game where id=?")

	if err != nil {
		panic(err.Error())
	}

	res, err := statement.Exec(strconv.Itoa(*id))

	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := res.RowsAffected()

	var payload model.MutateItemPayload

	if rowsAffected == 1 {
		payload = model.MutateItemPayloadSuccess{
			ID:      *id,
			Success: true,
			Message: "Game with id of " + strconv.Itoa(*id) + " has been deleted",
		}
	}

	if rowsAffected == 0 {
		payload = model.MutateItemPayloadFailure{
			ID:      *id,
			Success: false,
			Error:   "Delete operation not successful or did not exist.",
		}
	}

	return payload, nil
}
