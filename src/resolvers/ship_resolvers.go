package resolvers

import (
	"context"
	"errors"

	"github.com/jtomasevic/go-with-graphql-demo/src/graph/model"
	dataloaders "github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_loaders"
)

// shipResolver exists (generated from schema), because in gqlgen.yml we specified
// that we want to have resolver for some of ship fields. In our case this field is
// Ship.Crew. Example of configuration in gqlgen.yml:
//
// models:
//
//	   ...
//		  Ship:
//			fields:
//				crew:
//					resolver: true
type shipResolver struct{ *Resolver }

// Crew is the resolver for the Ship.Crew field. Every time when server receive request for
// ship with Crew, this method will be called to resolve Crew field. Let's say we have
// query: {ships {id, name, crew:{id, name}}}, and we have 10 ships in total. In this scenario
// shipResolver.Crew method will be called 10 times, for each ship.
//
// In resolvers like this, we  usually retreive data from some data source, and if we do this
// "n" time for "n" crews, without optimization we'll have (n+1) issue.
// Therefor, this is perfect place to use data loaders.
func (r *shipResolver) Crew(ctx context.Context, obj *model.Ship) (model.Crew, error) {
	// fmt.Printf("shipResolver->Crew, crew id:%s\n", obj.Crew.ID.String())

	if obj == nil {
		// if ship is nil, return empty struct, which should be recognized as nil on consumer side.
		// fo check if crew is empty we can check if Crew.ID == uuid.Nil
		return model.Crew{}, errors.New("for this resolver ship is mandatory")
	}
	// here we use data loader to prepare fetch statement and avoid (n+1) problem.
	crew, err := dataloaders.GetLoaders(ctx).CrewById.Load(obj.Crew.ID)
	if err != nil {
		return model.Crew{}, err
	}

	// just mapping to graph Model
	return crewFromServiceToGql(crew), nil
}
