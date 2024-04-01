// go:build wireinject
//go:build wireinject
// +build wireinject

package services

import (
	"github.com/google/wire"
	seas "github.com/jtomasevic/go-with-graphql-demo/src/services/7seas"
	datasource "github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_store/data_source"
	"github.com/jtomasevic/go-with-graphql-demo/src/services/7seas/data_store/repo"
)

// InilizeServices. wire is compile-time dependency injection library for Golang.
// First we link interface to it's implementation, and then provide factory method.
//
// Example:--------- interface------------------ implementation-----------
// wire.Bind(new(datasource.DataSource), new(*datasource.DataSourceService)),
// --- factory method.
// datasource.NewDataSource,
//
// Finally we how we bind return value. More about wire: https://github.com/google/wire
func InilizeServices() *Services {
	panic(wire.Build(

		wire.Bind(new(datasource.DataSource), new(*datasource.DataSourceService)),
		datasource.NewDataSource,

		wire.Bind(new(seas.DataStore), new(*repo.PiratesRepo)),
		seas.NewDataStore,

		wire.Bind(new(seas.Service), new(*seas.SevenSeasService)),
		seas.NewService,

		wire.Struct(new(Services), "*"),
	))
}
