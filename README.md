## About project 
This demo project demonstrates how to: 
1. Build a GraphQL project in Golang.
2. Demonstrate that Apollo Client can be integrated into a React application seamlessly, maintaining the same functionality as its current usage.

### Main topics
- Choosing framework.
  - Minimum requirements.
- Gqlgen
  - Schema first. 
  - Resolvers.
  - Field Resolvers.
    - (n+1) issue.
    - Data Loaders
## How does the GraphQL subgraph service, written in Golang integrate into the overall architecture?
The statement found, in the Apollo documentation, motivate us to explore an appropriate GraphQL solution in the Go language.
</br>
> Different subgraphs in the same supergraph can use different server implementations and even different programming languages as long as they are federation-compatible. 
> Source: [apollographql.com](https://www.apollographql.com/docs/federation/#next-steps)

**After all, it’s ”just” a protocol.** 

Multiple languages support the REST protocol, and even within the same language, there are multiple implementations. 

The same principle applies to GraphQL. As long as the server implement GrapQL protocol, it can be developed in any language.

#### However, there are some limitation.
- The chosen language must be suitable for a containerized environment; for instance, Python may not be optimal, despite its excellent performance in Lambda functions.

**Containerization environment**
- When deploying containerization solutions like ECS, Fargate, or EKS, we manage instances of servers that provide various APIs (in our case mostly REST and GraphQL). 
- We aim to avoid router logic, as in the container-based architecture, single node (server) is responsible for handling API requests, business logic, and resource access. 

## Choosing framework 

### Choices

There are several implementation of graphQL servers in Golang, like:

| Framework | Github stats |
| ----------- | ----------- |
| https://github.com/graphql-go/graphql | Stars: 9.6k; Forks: 870 |
| https://github.com/graph-gophers/graphql-go | Stars: 4.6k; Forks: 541 |
| https://github.com/samsarahq/thunder | Stars: 1.6k; Forks: 115 |
| https://github.com/wundergraph/graphql-go-tools | Stars: 615; Forks: 116 |

**Requested features**

Features that we are looking for are:
- Schema first approach (preferable with resolver generation).
- Easy implementation of data loaders (preferable with skeleton generation).
- Support for federation.
- Higley configurable. 

**Decision**

After investigating list above narrow us to [gqlgen](https://gqlgen.com/) Stars: 9.5k; Forks: 1.2k 

## gqlgen

**What is gqlgen**

By its own words:

> - [gqlgen](https://gqlgen.com/) a Go library for building GraphQL servers without any fuss.
> - gqlgen is based on a Schema first approach — You get to Define your API using the GraphQL [Schema Definition Language](http://graphql.org/learn/schema/).
> - gqlgen prioritizes Type safety 
> - gqlgen enables Codegen — We generate the boring bits, so you can focus on building your app quickly.

*To start new project with gqlgen follow https://gqlgen.com/getting-started/*

### Generation of models and resolvers code. ###

As gqlgen follows a schema-first approach framework, meaning that the starting point is the GraphQL schema. In our demo project this is schema:
```
scalar UUID

type Pirate {
  id: ID!
  name: String!
}

type Crew {
  id: ID!
  name: String!
  pirates: [Pirate!]!
}

type Ship {
  id: ID!
  name: String!
  crew: Crew!
}

input UpsertPirate {
    id: ID
    name: String!
    crewId: UUID
}

input UpsertCrew {
    id: ID
    name: String!
    shipId: UUID
}

input UpsertShip {
    id: ID
    name: String!
}

type Query {
  pirates: [Pirate!]!
  crews: [Crew!]!
  ships: [Ship!]!
  ship(id: UUID): Ship!
}

type Mutation {
  createPirate(input: UpsertPirate!): Pirate!
  createCrew(input: UpsertCrew!): Crew!
  createShip(input: UpsertShip!): Ship!
}
```

Running 
```
make generate
```
Will generate models, strongly typed interfaces for resolving queries and mutations, and their dummy implementations.

Models, interfaces, and specific graph code will be generated into:
- src/graph/
  - mode/models_gen.go
  - generated.go

**This code we do not touch**

Dummy resolvers implementation will be generated here:
- src/resolvers
  - schema.resolvers.go
  - resolver.go

> If we implement stubs in `schema.resolvers.go` file, and run the generate command again, the generator will not overwrite our changes.

We could keep everything into a single file, but this would likely result in a cumbersome and challenging-to-maintain document. To improve code readability, we'll divide the generated code into multiple files, assigning one for each resolver. Subsequently, we delete schema.resolver.go file. 

#### Example for generated resolvers:
**Schema**
```
  type Query {
    pirates: [Pirate!]!
    crews: [Crew!]!
    ships: [Ship!]!
    ship(id: UUID): Ship!
  }

  type Mutation {
    createPirate(input: UpsertPirate!): Pirate!
    createCrew(input: UpsertCrew!): Crew!
    createShip(input: UpsertShip!): Ship!
  }
  ```
  **Resolvers**
  ```
  type QueryResolver interface {
    Pirates(ctx context.Context) ([]model.Pirate, error)
    Crews(ctx context.Context) ([]model.Crew, error)
    Ships(ctx context.Context) ([]model.Ship, error)
    Ship(ctx context.Context, id *string) (model.Ship, error)
  }

  type MutationResolver interface {
    CreatePirate(ctx context.Context, input model.UpsertPirate) (model.Pirate, error)
    CreateCrew(ctx context.Context, input model.UpsertCrew) (model.Crew, error)
    CreateShip(ctx context.Context, input model.UpsertShip) (model.Ship, error)
  }
  ``` 

### Field resolvers. 
If we wish to resolve references of a particular type separately, we must first configure a field resolver in the gqlgen configuration file. The generator will then produce a new resolver interface for the referenced field(s).

**Models in graphQL schema**
```
type Pirate {
  id: ID!
  name: String!
}

type Crew {
  id: ID!
  name: String!
  pirates: [Pirate!]!
}

type Ship {
  id: ID!
  name: String!
  crew: Crew!
}
```

**Configuration in gqlgne.yaml**
```
models:
  ...
  Ship:
    fields:
      crew:
        resolver: true
  Crew:
    fields:
      pirates:
        resolver: true
```

**Generated resolver interface for Crew.Pirates and field Ship.Crew**
```
type CrewResolver interface {
	Pirates(ctx context.Context, obj *model.Crew) ([]model.Pirate, error)
}

type ShipResolver interface {
	Crew(ctx context.Context, obj *model.Ship) (model.Crew, error)
}
```
#### N+1 issue 

If we implement a field resolver to retrieve data from a data source, for example we’re easily jump into N+1 problem. For example If we have graphQL Query:

```
  query GetCrews {
    crews {
      id
      name
      pirates {
        id
        name
      }
    }
  }
```
We have one query for all crews:
```
SELECT * FROM crews
```
… and then for each crew (n times)
```
SELECT * FROM pirates where crew_id = ?
```
Our API would be much more efficient, if we can avoid N+1 issue. Let’s take a look to one more example:
![alt text](readme-images/4-queries.png)

### Data Loaders ###
Data loaders serve as a vital tool in addressing the N+1 problem. Their implementation typically dictates a specific approach to data retrieval, thereby influencing the design of service/repository APIs.

For this project, we are using the recommended library for generating data loaders, as found in gqlgen [examples](https://github.com/99designs/gqlgen/tree/master/_examples/dataloader): [github.com/vektah/dataloaden](https://github.com/vektah/dataloaden)

To generate data loaders run the command `make generate-dl`. (Chek this command to see how code generation works). This command will generate files in folder: `src/service/7seas/`, and then moved manually to `src/service/7seas/data_loaders`. 
- Gqlgen generate data loader skeleton, and then provide simple API for data loader definition.
- In the file `data_loaders.go` we implemented loaders basic logic and define middleware. 

This example explore two scenarios:
- Data loader for referenced object.
```
type Ship {
  id: ID!
  name: String!
  crew: Crew!
}
```

Ship has reference to Crew, and if we query multiple Ships without data loaders, we would have (n+1) problem.

- Data loader for child objects.
```
type Crew {
  id: ID!
  name: String!
  pirates: [Pirate!]!
}
```
Crew has array of pirates, and if we query multiple Crews without data loaders wa would also have (n+1) problem.

**Data loaders implementation**
```

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
```

**Using data loader: Crew resolver**
```
func (r *crewResolver) Pirates(ctx context.Context, obj *model.Crew) ([]model.Pirate, error) {

	pirates, err := dataloaders.GetLoaders(ctx).PiratesByCrews.Load(obj.ID)
	if err != nil {
		return nil, err
	}
	// just mapping to graph model.
	return piratesFromServiceToGql(pirates), nil
}
```
**Using data loader: Ship resolver**
```
func (r *shipResolver) Crew(ctx context.Context, obj *model.Ship) (model.Crew, error) {
	// here we use data loader to prepare fetch statement and avoid (n+1) problem.
	crew, err := dataloaders.GetLoaders(ctx).CrewById.Load(obj.Crew.ID)
	if err != nil {
		return model.Crew{}, err
	}
	// just mapping to graph Model
	return crewFromServiceToGql(crew), nil
}
```

## Run example ###

```
make run 
```

Open `http://localhost:8080/graphql` and here are example of query you can try:

```
{
	ships{
    id, 
    name,
    crew{
      id
      name
      pirates{
        id
        name
      }
    }
  }
}
```

> Check the terminal to observe the order in which resolvers are called, and the SQL commands that are executed.
