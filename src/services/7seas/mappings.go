package seas

import "github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_store/model"

func pirateFromDbToService(row model.Pirate) Pirate {
	return Pirate{
		ID:     row.Id,
		Name:   row.Name,
		CrewId: *row.CrewId,
	}
}

func piratesFromDbToService(rows []model.Pirate) []Pirate {
	results := []Pirate{}
	for _, row := range rows {
		results = append(results, pirateFromDbToService(row))
	}
	return results
}

func crewFromDbToService(row model.Crew) Crew {
	return Crew{
		ID:   row.Id,
		Name: row.Name,
	}
}

func crewsFromDbToService(rows []model.Crew) []Crew {
	results := []Crew{}
	for _, row := range rows {
		results = append(results, crewFromDbToService(row))
	}
	return results
}

func shipFromDbToService(row model.Ship) Ship {
	return Ship{
		ID:   row.Id,
		Name: row.Name,
		Crew: &Crew{
			ID: *row.CrewId,
		},
	}
}

func shipsFromDbToService(rows []model.Ship) []Ship {
	results := []Ship{}
	for _, row := range rows {
		results = append(results, shipFromDbToService(row))
	}
	return results
}
