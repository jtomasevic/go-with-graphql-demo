package resolvers

import (
	"context"
	"fmt"

	"github.com/jtomasevic/go-with-graphql-demo/src/graph/model"
)

type mutationResolver struct {
	*Resolver
}

// CreatePirate is the resolver for the createPirate field.
func (r *mutationResolver) CreatePirate(ctx context.Context, input model.UpsertPirate) (model.Pirate, error) {
	panic(fmt.Errorf("not implemented: CreatePirate - createPirate"))
}

// CreateCrew is the resolver for the createCrew field.
func (r *mutationResolver) CreateCrew(ctx context.Context, input model.UpsertCrew) (model.Crew, error) {
	panic(fmt.Errorf("not implemented: CreateCrew - createCrew"))
}

// CreateShip is the resolver for the createShip field.
func (r *mutationResolver) CreateShip(ctx context.Context, input model.UpsertShip) (model.Ship, error) {
	panic(fmt.Errorf("not implemented: CreateShip - createShip"))
}
