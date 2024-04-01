package resolvers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jtomasevic/go-with-graphql-demo/src/server"
	"github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_store/repo"
	"github.com/stretchr/testify/require"
)

const (
	shipsQuery   = "{ships {id, name, crew { id, name, pirates { id, name }}}}"
	crewQuery    = "{crews {id, name, pirates{ id, name }}}"
	piratesQuery = "{pirates { id, name }}"
)

func BenchmarkLoad(b *testing.B) {
	handler := server.NewGQLHttpHandler()
	// 2 ships, 2 crews, 16 pirates
	b.Run("pirates", func(b *testing.B) {
		runRequest(b, handler, piratesQuery)
	})
	b.Run("crews", func(b *testing.B) {
		runRequest(b, handler, crewQuery)
	})
	b.Run("ships", func(b *testing.B) {
		runRequest(b, handler, shipsQuery)
	})

	err := repo.PopulateDb(true)
	require.NoError(b, err)
	// 200 ships, 200 crews, 200 pirates per crew = 40.000
	b.Run("pirates with large data set", func(b *testing.B) {
		runRequest(b, handler, piratesQuery)
	})
	b.Run("crews with large data set", func(b *testing.B) {
		runRequest(b, handler, crewQuery)
	})
	b.Run("ships with large data set", func(b *testing.B) {
		runRequest(b, handler, shipsQuery)
	})
}

func runRequest(B *testing.B, r http.Handler, query string) {
	req, err := http.NewRequest("POST", "/graphql", strings.NewReader(
		fmt.Sprintf("{\"query\": \"%s\", \"operationName\":\"\", \"variables\": null}", query)))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	B.ReportAllocs()
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		r.ServeHTTP(w, req)
	}
}
