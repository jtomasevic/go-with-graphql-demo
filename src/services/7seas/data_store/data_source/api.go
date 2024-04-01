package datasource

import (
	"database/sql"
	"log"

	_ "github.com/proullon/ramsql/driver"
)
type DataSource interface {
	Db()  *sql.DB
}

func NewDataSource() *DataSourceService {
	db, err := sql.Open("ramsql", "Pirates")
	if err != nil {
		log.Fatalf("sql.Open : Error : %s\n", err)
		panic(err)
	}

	return &DataSourceService{
		db: db,
	}
}