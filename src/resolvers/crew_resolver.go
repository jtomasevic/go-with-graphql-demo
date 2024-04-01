package resolvers

import (
	"context"

	"github.com/jtomasevic/go-with-graphql-demo/src/graph/model"
	dataloaders "github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_loaders"
)

// crewResolver exists (generated from schema), because in gqlgen.yml we specified
// that we want to have resolver for some of crew fields. In our case this field is
// Crew.Pirates. Example of configuration in gqlgen.yml:
//
// models:
//
//	  ....
//	  Crew:
//		fields:
//			pirates:
//				resolver: true
type crewResolver struct {
	*Resolver
}

// Pirates is the resolver for the Crew.Pirates field. Every time when server receive request for
// crew with Pirates, this method will be called to resolve Pirates collection. Let's say we have
// query: {crews {id, name, pirates:{id, name}}}, and we have 10 crews in total.
//
// In resolvers like this, we usually retreive data from some data source, and if we do this
// "n" time for "n" crews, without optimization we'll have (n+1) issue.
// Therefor, this is perfect place to use data loaders.
func (r *crewResolver) Pirates(ctx context.Context, obj *model.Crew) ([]model.Pirate, error) {

	// fmt.Printf("crewResolver->Pirates, crew id:%s\n", obj.ID.String())
	if obj == nil {
		// if crew is nil, then there just return empty collection.
		return []model.Pirate{}, nil
	}
	// here we use data loader to prepare fetch statement and avoid (n+1) problem.
	pirates, err := dataloaders.GetLoaders(ctx).PiratesByCrews.Load(obj.ID)
	if err != nil {
		return nil, err
	}
	// just mapping to graph model.
	return piratesFromServiceToGql(pirates), nil
}
