package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/davidfunk13/overwatch-companion/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver : holds model references
type Resolver struct {
	battletags []*model.Battletag
	sessions []model.Session
	games []model.Game
}
