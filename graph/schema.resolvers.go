package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/davidfunk13/overwatch-companion/database"
	"github.com/davidfunk13/overwatch-companion/graph/generated"
	"github.com/davidfunk13/overwatch-companion/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.InputUser) (*model.User, error) {
	user := &model.User{
		Name:  input.Name,
		Email: input.Email,
	}

	r.users = append(r.users, user)

	res, err := database.Db.Exec(`INSERT INTO user (name, email) VALUES (?, ?);`, user.Name, user.Email)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Successfully inserted")
	}

	lastInsertedID, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	lastInserted := database.Db.QueryRow(`Select * from user where id=?;`, lastInsertedID)

	if err != nil {
		panic(err.Error())
	}

	var id, name, email string

	err = lastInserted.Scan(&id, &name, &email)

	insertedUser := model.User{ID: id, Name: name, Email: email}

	return &insertedUser, nil
}

func (r *mutationResolver) CreateBattletag(ctx context.Context, input model.InputBattletag) (*model.Battletag, error) {
	battletagInput := &model.Battletag{
		UserID:     input.UserID,
		Battletag:  input.Battletag,
		Platform:   &input.Platform,
		Identifier: input.Identifier,
	}

	r.battletags = append(r.battletags, battletagInput)

	res, err := database.Db.Exec(`INSERT INTO battletags (userId, battletag, platform, identifier) VALUES (?, ?, ?, ?);`, battletagInput.UserID, battletagInput.Battletag, battletagInput.Platform, battletagInput.Identifier)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Successfully inserted")
	}

	lastInsertedID, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	lastInserted := database.Db.QueryRow(`Select * from battletag where id=?;`, lastInsertedID)

	if err != nil {
		panic(err.Error())
	}

	var id, userId, battletag string
	var identifier *string
	var platform *model.Platform

	err = lastInserted.Scan(&id, &userId, &battletag, &platform, &identifier)
	insertedBattletag := model.Battletag{ID: id, UserID: userId, Battletag: battletag, Platform: platform, Identifier: identifier}

	return &insertedBattletag, nil
}

func (r *mutationResolver) DeleteBattletag(ctx context.Context, input string) (model.MutateBattletagPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var data = database.SelectAllUsers()

	return data, nil
}

func (r *queryResolver) Battletags(ctx context.Context) ([]*model.Battletag, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
