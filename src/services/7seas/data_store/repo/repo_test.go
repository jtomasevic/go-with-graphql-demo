package repo

import (
	"context"
	"testing"

	"github.com/google/uuid"
	datasource "github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_store/data_source"
	"github.com/stretchr/testify/require"
)

func TestInitDb(t *testing.T) {
	err := InitDb()
	require.NoError(t, err)
}
func TestPopulateDb(t *testing.T) {
	initDb(t)

}

func TestGetAll(t *testing.T) {
	initDb(t)
	repo := PiratesRepo{
		DataSource: datasource.NewDataSource(),
	}
	// pirates
	t.Run("get all pirates.", func(t *testing.T) {
		pirates, err := repo.GetPirates(context.TODO(), nil)
		require.NoError(t, err)
		require.NotEmpty(t, pirates)
	})
	t.Run("get all pirates for some crew.", func(t *testing.T) {
		pirates, err := repo.GetPirates(context.TODO(), &cursedCrewId)
		require.NoError(t, err)
		require.NotEmpty(t, pirates)
		for _, pirate := range pirates {
			require.Equal(t, *pirate.CrewId, cursedCrewId)
		}
	})
	// crews
	t.Run("get all crews.", func(t *testing.T) {
		crews, err := repo.GetCrews(context.TODO())
		require.NoError(t, err)
		require.NotEmpty(t, crews)
	})

	// ships
	t.Run("get all ships.", func(t *testing.T) {
		ships, err := repo.GetShips(context.TODO())
		require.NoError(t, err)
		require.NotEmpty(t, ships)
	})
}

func TestGetByIds(t *testing.T) {
	initDb(t)
	t.Run("crews", func(t *testing.T) {
		repo := PiratesRepo{
			DataSource: datasource.NewDataSource(),
		}
		crews, err := repo.GetCrews(context.TODO())
		require.NoError(t, err)
		require.NotEmpty(t, crews)

		ids := []uuid.UUID{}
		for _, crew := range crews {
			ids = append(ids, crew.Id)
		}

		byIdsCrews, err := repo.GetCrewsByIds(context.TODO(), ids)
		require.NoError(t, err)
		require.NotEmpty(t, byIdsCrews)

		require.Equal(t, len(crews), len(byIdsCrews))
	})

	t.Run("pirates", func(t *testing.T) {
		repo := PiratesRepo{
			DataSource: datasource.NewDataSource(),
		}
		pirates, err := repo.GetPirates(context.TODO(), nil)
		require.NoError(t, err)
		require.NotEmpty(t, pirates)

		ids := []uuid.UUID{}
		for _, pirate := range pirates {
			ids = append(ids, pirate.Id)
		}

		byIdsPirates, err := repo.GetPiratesByIds(context.TODO(), ids)
		require.NoError(t, err)
		require.NotEmpty(t, byIdsPirates)

		require.Equal(t, len(pirates), len(byIdsPirates))
	})
}

func TestByFK(t *testing.T) {
	initDb(t)
	t.Run("Get crew by ship id", func(t *testing.T) {
		repo := PiratesRepo{
			DataSource: datasource.NewDataSource(),
		}
		crew, err := repo.GetCrewForShip(context.TODO(), blackPearlId)
		require.NoError(t, err)
		require.Equal(t, crew.Id, cursedCrewId)
	})

	t.Run("Get crew by ship id", func(t *testing.T) {
		repo := PiratesRepo{
			DataSource: datasource.NewDataSource(),
		}
		pirates, err := repo.GetPiratesInCrews(context.TODO(), []uuid.UUID{cursedCrewId, flyingDutchmanCrewId})
		require.NoError(t, err)
		require.NotEmpty(t, pirates)
	})
}

func TestById(t *testing.T) {
	initDb(t)
	t.Run("Get ship by id", func(t *testing.T) {
		repo := PiratesRepo{
			DataSource: datasource.NewDataSource(),
		}
		ships, err := repo.GetShips(context.TODO())
		require.NoError(t, err)
		require.NotEmpty(t, ships)

		ship, err := repo.GetShip(context.TODO(), ships[0].Id)
		require.NoError(t, err)
		require.EqualValues(t, ships[0], ship)
	})

	t.Run("Get crew by id", func(t *testing.T) {
		repo := PiratesRepo{
			DataSource: datasource.NewDataSource(),
		}
		crewa, err := repo.GetCrews(context.TODO())
		require.NoError(t, err)
		require.NotEmpty(t, crewa)

		crew, err := repo.GetCrew(context.TODO(), crewa[0].Id)
		require.NoError(t, err)
		require.EqualValues(t, crewa[0], crew)
	})
}

func initDb(t *testing.T) {

	err := InitDb()
	require.NoError(t, err)
	err = PopulateDb(false)
	require.NoError(t, err)
}
