package databaselayer

import (
	"database/sql"
	// This blank import is necessary to execute the init package function and register the "sqlite" driver.
	_ "github.com/mattn/go-sqlite3"
)

// SQLiteHandler is the handler to use with SQLite databases, it is an object that encapsulates the generic SQLHandler using "composition".
type SQLiteHandler struct {
	*SQLHandler
}

// NewSQLiteHandler is the constructor to build a handler for the SQLite dbtype
func NewSQLiteHandler(connection string) (*SQLiteHandler, error) {
	db, err := sql.Open("sqlite3", connection)
	return &SQLiteHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
