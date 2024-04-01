package resolvers

import (
	"github.com/jtomasevic/go-with-graphql-demo/src/graph"
	"github.com/jtomasevic/go-with-graphql-demo/src/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Services services.Services
}


// Crew returns graph.CrewResolver implementation.
func (r *Resolver) Crew() graph.CrewResolver { 
	return &crewResolver{r} 
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { 
	return &mutationResolver{r} 
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { 
	return &queryResolver{r} 
}

// Ship returns graph.ShipResolver implementation.
func (r *Resolver) Ship() graph.ShipResolver { 
	return &shipResolver{r} 
}