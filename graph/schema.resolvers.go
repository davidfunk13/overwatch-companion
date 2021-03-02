package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/davidfunk13/overwatch-companion/database"
	"github.com/davidfunk13/overwatch-companion/graph/generated"
	"github.com/davidfunk13/overwatch-companion/graph/model"
)

func (r *mutationResolver) CreateBattletag(ctx context.Context, input model.InputBattletag) (*model.Battletag, error) {
	inserted := database.CreateBattletag(input)

	//why do we do this?!? why not just create and return; query and return? why append to resolver?
	r.battletags = append(r.battletags, inserted)

	return inserted, nil
}

func (r *mutationResolver) DeleteBattletag(ctx context.Context, input int) (model.MutateBattletagPayload, error) {
	// database.DeleteBattletag(&input)

	payload := model.BattletagMutationSuccess{
		ID:      1,
		Success: true,
		Message: "Successfully deleted",
	}

	return payload, nil
}

func (r *queryResolver) Battletags(ctx context.Context) ([]*model.Battletag, error) {
	var data = database.SelectAllBattletags()

	return data, nil

	// or this?!?!??!
	// return r.battletags, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
