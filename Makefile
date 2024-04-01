PHONY: init setup generate install generate-dl


# this command is run only for the first time to install and setup gqlgen 
# library and add gqlgen as a tool dependency for your module.
# !!!! DON'T RUN THIS COMMAND IF PROJECT IS ALREADY SETUP !!!
setup:
	# Add github.com/99designs/gqlgen to your project's tools.go
	printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
	go mod tidy

# this command is run only for the first time to generate code, as it's
# define in gqlgen.yaml
init:
	go run github.com/99designs/gqlgen init

# open http://localhost:8080/graphql for local playground. 
generate:
	go run github.com/99designs/gqlgen generate
	
install: 
	go get github.com/proullon/ramsql
	go get github.com/google/uuid
	go get github.com/stretchr/testify/require
	go get github.com/vektah/dataloaden
	go get github.com/google/wire
	go get github.com/rs/cors

# generate data loaders
# cd src/services/7seas; 
generate-dl:
	cd src/services/7seas; \
		go run github.com/vektah/dataloaden CrewLoader github.com/google/uuid.UUID 'github.com/jtomasevic/go-with-graphql-demo/src/services/seas.Crew'; \
		go run github.com/vektah/dataloaden PiratesLoader github.com/google/uuid.UUID '[]gitlab.com/jovantomasevic/gql-example/src/services/pirates.Pirate'

run:
	go run server.go

