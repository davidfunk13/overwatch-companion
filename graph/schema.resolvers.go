package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/davidfunk13/overwatch-companion/graph/generated"
	"github.com/davidfunk13/overwatch-companion/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	//create empty variable modeling after todo
	var todo model.Todo
	//create empty variable modeling after user
	var user model.User
	//set Text and Done in todo to sample data
	todo.Text = input.Text
	todo.ID = input.UserID
	// set user on new todo to a pointer to the address of the user instantiated above
	todo.User = &user
	// return a pointer to the address of
	// the instantiated todo object above.
	return &todo, nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// define an empty variable with a type of Todo from our Models
	var todos []*model.Todo
	// create a new todo by instantiating a new Todo with this syntax
	dummyTodo := model.Todo{
		ID:   "1",
		Text: "Finish an application",
		Done: false,
		User: &model.User{Name: "David"},
	}

	todos = append(todos, &dummyTodo)

	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
