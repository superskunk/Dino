package databaselayer

import (
	"database/sql"
	// This blank import is necessary to execute the package init function to register the "mysql" drvier according to "datbase/sql".
	_ "github.com/go-sql-driver/mysql"
)

// MySQLHandler is the handler to use with MySQL databases, it is an object that encapsulates the generic SQLHandler using "composition".
type MySQLHandler struct {
	*SQLHandler
}

// NewMySQLHandler is the constructor to build a handler for the MySQL dbtype
func NewMySQLHandler(connection string) (*MySQLHandler, error) {
	db, err := sql.Open("mysql", connection)
	return &MySQLHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
