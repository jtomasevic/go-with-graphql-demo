package seas

import (
	"context"

	"github.com/google/uuid"
	datasource "github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_store/data_source"
	"github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_store/model"
	"github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_store/repo"
)

// service interface
type Service interface {
	GetPirates(ctx context.Context, crew_id *uuid.UUID) ([]Pirate, error)
	GetPiratesInCrews(ctx context.Context, crew_ids []uuid.UUID) ([]Pirate, error)
	GetPiratesByIds(ctx context.Context, ids []uuid.UUID) ([]Pirate, error)
	GetCrews(ctx context.Context) ([]Crew, error)
	GetCrew(ctx context.Context, id uuid.UUID) (Crew, error)
	GetCrewsByIds(ctx context.Context, ids []uuid.UUID) ([]Crew, error)
	GetCrewForShip(ctx context.Context, shipId uuid.UUID) (Crew, error)
	GetShips(ctx context.Context) ([]Ship, error)
	GetShip(ctx context.Context, id uuid.UUID) (Ship, error)
}

// service api also define interface for data access. we are using here pattern where consumer define interfaces for producers.
// mora about this pattern:
type DataStore interface {
	GetPirates(ctx context.Context, crew_id *uuid.UUID) ([]model.Pirate, error)
	GetPiratesInCrews(ctx context.Context, crew_ids []uuid.UUID) ([]model.Pirate, error)
	GetPiratesByIds(ctx context.Context, ids []uuid.UUID) ([]model.Pirate, error)
	GetCrews(ctx context.Context) ([]model.Crew, error)
	GetCrew(ctx context.Context, id uuid.UUID) (model.Crew, error)
	GetCrewForShip(ctx context.Context, shipId uuid.UUID) (model.Crew, error)
	GetCrewsByIds(ctx context.Context, ids []uuid.UUID) ([]model.Crew, error)
	GetShips(ctx context.Context) ([]model.Ship, error)
	GetShip(ctx context.Context, id uuid.UUID) (model.Ship, error)
}

func NewDataStore(dataSource datasource.DataSource) *repo.PiratesRepo {
	repo.InitDb()
	repo.PopulateDb(false)
	return &repo.PiratesRepo{
		DataSource: dataSource,
	}
}

func NewService(dataStore DataStore) *SevenSeasService {
	return &SevenSeasService{
		dataStore: dataStore,
	}
}
