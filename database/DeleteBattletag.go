package database

//DeleteBattletag deletes a single battletag from the database by id
// func DeleteBattletag(id *int) (model.MutateBattletagPayload, error) {
// var success bool
// var message string

// var userId int
// var identifier *int
// var battletag string
// var platform *model.Platform
// res := Db.QueryRow("DELETE FROM battletag where id=?", id)

// row := res.Scan(&id, &userId, &battletag, &platform, &identifier)

// payload := model.BattletagMutationSuccess{
// 	ID:      1,
// 	Success: true,
// 	Message: "Successfully deleted",
// }

// // payload := res.Scan()
// return payload, nil

// }
