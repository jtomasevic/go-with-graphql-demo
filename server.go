package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jtomasevic/go-with-graphql-demo/src/server"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// create http graphql handler
	handler := server.NewGQLHttpHandler()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// handle graphQL request
	http.Handle("/query", handler)

	// if we need to serve REST calls we would do something like this:
	// mux := http.NewServeMux()
	// mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		w.Write([]byte("Hello!"))
	// 	}
	// })

	// handle other requests as REST calls
	// http.Handle("/hello", mux)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
