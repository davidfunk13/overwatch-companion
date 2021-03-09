package database

import (
	"strconv"

	"github.com/davidfunk13/overwatch-companion/graph/model"
)

//DeleteBattletag deletes a single battletag from the database by id
func DeleteBattletag(id *int) (model.MutateItemPayload, error) {
	db, err := OpenConnection()

	defer db.Close()
	
	//prepare statement to delete the battletag itself.
	statement, err := db.Prepare("DELETE FROM battletag where id=?")

	if err != nil {
		panic(err.Error())
	}

	//delete the battletag
	res, err := statement.Exec(strconv.Itoa(*id))

	if err != nil {
		panic(err.Error())
	}

	//lets us know if the delete was succeessful.
	rowsAffected, err := res.RowsAffected()

	var payload model.MutateItemPayload

	if rowsAffected == 1 {
		payload = model.MutateItemPayloadSuccess{
			ID:      *id,
			Success: true,
			Message: "Battletag with id of " + strconv.Itoa(*id) + " has been deleted",
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
