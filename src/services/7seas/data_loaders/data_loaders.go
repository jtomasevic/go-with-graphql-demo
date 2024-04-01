package dataloaders

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	seas "github.com/jtomasevic/go-with-graphql-demo/src/services/7seas"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"userCtx"}

type Loaders struct {
	CrewById       *CrewLoader
	PiratesByCrews *PiratesLoader
}

func NewLoaders(ctx context.Context, service seas.Service) (Loaders, context.Context) {
	dataLoaders := Loaders{}

	// how long to done before sending a batch
	wait := 150 * time.Microsecond

	dataLoaders.CrewById = &CrewLoader{
		wait:     wait,
		maxBatch: 1000,
		fetch: func(keys []uuid.UUID) ([]seas.Crew, []error) {
			crews, err := service.GetCrewsByIds(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}
			// some data bases doesn't return values in expected order. i.e if in SQL statemente we have
			// ... WHERE some_id IN ('abc', 'xyz') this doesn't necessary means that first row in result
			// will be 'abc' and second 'xyz'.
			// So we need to make sure we'll return list of crews in the order defined by input param @keys
			result := make([]seas.Crew, len(crews))
			crewMap := toCrewsMap(crews)
			for i, key := range keys {
				result[i] = crewMap[key]
			}
			return result, nil
		},
	}

	dataLoaders.PiratesByCrews = &PiratesLoader{
		wait:     wait,
		maxBatch: 2000,
		fetch: func(keys []uuid.UUID) ([][]seas.Pirate, []error) {
			pirates, err := service.GetPiratesInCrews(ctx, keys)
			result := make([][]seas.Pirate, len(keys))
			if err != nil {
				return nil, []error{err}
			}
			piratesMap := toPiratesMap(pirates)
			for i, key := range keys {
				result[i] = piratesMap[key]
			}
			return result, nil
		},
	}

	dataLoadersContext := context.WithValue(ctx, ctxKey, dataLoaders)
	return dataLoaders, dataLoadersContext
}

func LoaderMiddleware(next http.Handler, service seas.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ctx := NewLoaders(r.Context(), service)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetLoaders(ctx context.Context) Loaders {
	return ctx.Value(ctxKey).(Loaders)
}

func toCrewsMap(crews []seas.Crew) map[uuid.UUID]seas.Crew {
	dataMap := make(map[uuid.UUID]seas.Crew)
	for _, crew := range crews {
		dataMap[crew.ID] = crew
	}
	return dataMap
}

// toPiratesMap: in return result key is crew id, and value is list of pirates
func toPiratesMap(pirates []seas.Pirate) map[uuid.UUID][]seas.Pirate {
	dataMap := make(map[uuid.UUID][]seas.Pirate)
	for _, pirate := range pirates {
		dataMap[pirate.CrewId] = append(dataMap[pirate.CrewId], pirate)
	}
	return dataMap
}
