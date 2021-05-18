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

func (r *queryResolver) GetAllBattletags(ctx context.Context, input string) ([]*model.Battletag, error) {
	battletags := database.GetAllBattletags(input)

	return battletags, nil
}

func (r *queryResolver) GetOneBattletag(ctx context.Context, input *model.InputGetOneBattletag) ([]*model.Battletag, error) {
	battletag := database.GetOneBattletag(input)

	return battletag, nil
}

func (r *queryResolver) GetAllSessions(ctx context.Context, input *model.InputGetSessions) ([]*model.Session, error) {
	sessions := database.GetAllSessions(input.UserID, input.BattletagID)

	return sessions, nil
}

func (r *queryResolver) GetOneSession(ctx context.Context, input *model.InputGetOneSession) ([]*model.Session, error) {
	session := database.GetOneSession(input)

	return session, nil
}

func (r *queryResolver) GetMostRecentSession(ctx context.Context, input *model.InputGetMostRecentSession) ([]*model.Session, error) {
	session := database.GetMostRecentSession(input)

	return session, nil
}

func (r *queryResolver) GetAllGames(ctx context.Context, input *model.InputGetGames) ([]*model.Game, error) {
	games := database.GetAllGames(*input)

	return games, nil
}

func (r *queryResolver) GetOneGame(ctx context.Context, input *model.InputGetGame) ([]*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
