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

func (r *mutationResolver) CreateBattletag(ctx context.Context, input model.InputBattletag) (model.MutateItemPayload, error) {
	inserted := database.CreateBattletag(input)

	return inserted, nil
}

func (r *mutationResolver) DeleteBattletag(ctx context.Context, input int) (model.MutateItemPayload, error) {
	deleted, err := database.DeleteBattletag(&input)

	if err != nil {
		panic(err.Error())
	}

	return deleted, nil
}

func (r *mutationResolver) CreateSession(ctx context.Context, input model.InputSession) (model.MutateItemPayload, error) {
	inserted := database.CreateSession(input)

	//why do we do this?!? why not just create and return; query and return? why append to resolver?
	// r.sessions = append(r.sessions, inserted)

	return inserted, nil
}

func (r *mutationResolver) UpdateSessionStartingSr(ctx context.Context, input model.InputUpdateSessionStartingSr) (model.MutateItemPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSession(ctx context.Context, input int) (model.MutateItemPayload, error) {
	deleted, err := database.DeleteSession(&input)

	if err != nil {
		panic(err.Error())
	}

	return deleted, nil
}

func (r *mutationResolver) CreateGame(ctx context.Context, input model.InputGame) (model.MutateItemPayload, error) {
	inserted := database.CreateGame(input)

	//why do we do this?!? why not just create and return; query and return? why append to resolver?
	// r.games = append(r.games, inserted)

	return inserted, nil
}

func (r *mutationResolver) DeleteGame(ctx context.Context, input int) (model.MutateItemPayload, error) {
	deleted, err := database.DeleteSession(&input)

	if err != nil {
		panic(err.Error())
	}

	return deleted, nil
}

func (r *queryResolver) Battletags(ctx context.Context, input int) ([]*model.Battletag, error) {
	battletags, err := database.SelectAllBattletags(input)

	if err != nil {
		panic(err.Error())
	}

	return battletags, nil
}

func (r *queryResolver) Sessions(ctx context.Context, input *model.InputGetSessions) ([]*model.Session, error) {
	sessions, err := database.SelectAllSessions(input.UserID, input.BattletagID)

	if err != nil {
		panic(err.Error())
	}

	return sessions, nil
}

func (r *queryResolver) Games(ctx context.Context, input *model.InputGetGames) ([]*model.Game, error) {
	games, err := database.SelectAllGames(*input)

	if err != nil {
		panic(err.Error())
	}

	return games, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) UpdateStartingSr(ctx context.Context, input model.InputSession) (model.MutateItemPayload, error) {
	panic(fmt.Errorf("not implemented"))
}
