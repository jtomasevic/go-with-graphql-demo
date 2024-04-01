package seas_test

import (
	"context"
	"testing"

	"github.com/jtomasevic/go-with-graphql-demo/src/services"
	_ "github.com/proullon/ramsql/driver"
	"github.com/stretchr/testify/require"
)

func TestGetPirates(t *testing.T) {
	service := services.InilizeServices().SevenSeasService

	pirates, err := service.GetPirates(context.TODO(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, pirates)
}

func TestGetCrews(t *testing.T) {
	service := services.InilizeServices().SevenSeasService

	crews, err := service.GetCrews(context.TODO())
	require.NoError(t, err)
	require.NotEmpty(t, crews)
}

func TestGetShips(t *testing.T) {
	service := services.InilizeServices().SevenSeasService

	ships, err := service.GetShips(context.TODO())
	require.NoError(t, err)
	require.NotEmpty(t, ships)
}
