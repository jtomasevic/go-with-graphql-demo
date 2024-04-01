package datasource

import (
	"database/sql"

	_ "github.com/proullon/ramsql/driver"
)

type DataSourceService struct {
	db *sql.DB
}

func (dataSource *DataSourceService) Db() *sql.DB{
	return dataSource.db
}