package resolvers

import (
	"github.com/jtomasevic/go-with-graphql-demo/src/graph/model"
	seas "github.com/jtomasevic/go-with-graphql-demo/src/services/7seas"
)

func piratesFromServiceToGql(pirates []seas.Pirate) []model.Pirate {
	result := []model.Pirate{}
	for _, pirate := range pirates {
		result = append(result, model.Pirate{
			ID:   pirate.ID,
			Name: pirate.Name,
		})
	}
	return result
}

func crewFromServiceToGql(crew seas.Crew) model.Crew {
	return model.Crew{
		ID:      crew.ID,
		Name:    crew.Name,
		Pirates: piratesFromServiceToGql(crew.Pirates),
	}
}

func crewsFromServiceToGql(crews []seas.Crew) []model.Crew {
	result := []model.Crew{}
	for _, crew := range crews {
		result = append(result, crewFromServiceToGql(crew))
	}
	return result
}

func shipsFromServiceToGql(ships []seas.Ship) []model.Ship {
	result := []model.Ship{}
	for _, ship := range ships {
		result = append(result, shipFromServiceToGql(ship))
	}
	return result
}

func shipFromServiceToGql(ship seas.Ship) model.Ship {
	return model.Ship{
		ID:   ship.ID,
		Name: ship.Name,
		Crew: &model.Crew{
			ID: ship.Crew.ID,
		},
	}
}
