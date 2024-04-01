package resolvers

import (
	"context"

	"github.com/google/uuid"
	"github.com/jtomasevic/go-with-graphql-demo/src/graph/model"
)

type queryResolver struct {
	*Resolver
}

// Pirates is the resolver for the pirates field.
func (r *queryResolver) Pirates(ctx context.Context) ([]model.Pirate, error) {
	//fmt.Println("queryResolver->Pirates")
	pirates, err := r.Services.SevenSeasService.GetPirates(ctx, nil)
	if err != nil {
		return nil, err
	}
	return piratesFromServiceToGql(pirates), nil
}

// Crews is the resolver for the crews field.
func (r *queryResolver) Crews(ctx context.Context) ([]model.Crew, error) {
	//fmt.Println("queryResolver->Crews")
	crews, err := r.Services.SevenSeasService.GetCrews(ctx)
	if err != nil {
		return nil, err
	}
	return crewsFromServiceToGql(crews), nil
}

// Ships is the resolver for the ships field.
func (r *queryResolver) Ships(ctx context.Context) ([]model.Ship, error) {
	//fmt.Println("queryResolver->Ships")
	ships, err := r.Services.SevenSeasService.GetShips(ctx)
	if err != nil {
		return nil, err
	}
	return shipsFromServiceToGql(ships), nil
}

func (r *queryResolver) Ship(ctx context.Context, id *string) (model.Ship, error) {
	ship, err := r.Services.SevenSeasService.GetShip(ctx, uuid.MustParse(*id))
	if err != nil {
		return model.Ship{}, err
	}
	return shipFromServiceToGql(ship), nil
}
