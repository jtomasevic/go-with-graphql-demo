package seas

import (
	"context"

	"github.com/google/uuid"
)

type SevenSeasService struct {
	dataStore DataStore
}

func (service *SevenSeasService) GetPirates(ctx context.Context, crew_id *uuid.UUID) ([]Pirate, error) {
	fromDb, err := service.dataStore.GetPirates(ctx, crew_id)
	if err != nil {
		return nil, err
	}
	return piratesFromDbToService(fromDb), nil
}

func (service *SevenSeasService) GetPiratesByIds(ctx context.Context, ids []uuid.UUID) ([]Pirate, error) {
	fromDb, err := service.dataStore.GetPiratesByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	pirates := piratesFromDbToService(fromDb)
	return pirates, nil
}

func (service *SevenSeasService) GetPiratesInCrews(ctx context.Context, crew_ids []uuid.UUID) ([]Pirate, error) {
	fromDb, err := service.dataStore.GetPiratesInCrews(ctx, crew_ids)
	if err != nil {
		return nil, err
	}
	pirates := piratesFromDbToService(fromDb)
	return pirates, nil
}

func (service *SevenSeasService) GetCrews(ctx context.Context) ([]Crew, error) {
	fromDb, err := service.dataStore.GetCrews(ctx)
	if err != nil {
		return nil, err
	}
	crews := crewsFromDbToService(fromDb)
	return crews, nil
}

func (service *SevenSeasService) GetCrew(ctx context.Context, id uuid.UUID) (Crew, error) {
	fromDb, err := service.dataStore.GetCrew(ctx, id)
	if err != nil {
		return Crew{}, err
	}
	crew := crewFromDbToService(fromDb)
	return crew, nil
}

func (service *SevenSeasService) GetCrewForShip(ctx context.Context, shipId uuid.UUID) (Crew, error) {
	fromDb, err := service.dataStore.GetCrewForShip(ctx, shipId)
	if err != nil {
		return Crew{}, err
	}
	return crewFromDbToService(fromDb), nil 
}

func (service *SevenSeasService) GetCrewsByIds(ctx context.Context, ids []uuid.UUID) ([]Crew, error) {
	fromDb, err := service.dataStore.GetCrewsByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	crews := crewsFromDbToService(fromDb)
	return crews, nil
}

func (service *SevenSeasService) GetShips(ctx context.Context) ([]Ship, error) {
	fromDb, err := service.dataStore.GetShips(ctx)
	if err != nil {
		return nil, err
	}
	ships := shipsFromDbToService(fromDb)
	return ships, nil
}

func (service *SevenSeasService) GetShip (ctx context.Context, id uuid.UUID) (Ship, error) {
	fromDb, err := service.dataStore.GetShip(ctx, id)
	if err != nil {
		return Ship{}, err
	}
	ship := shipFromDbToService(fromDb)
	return ship, nil
}