package databaselayer

import (
	"database/sql"

	// This blank import is necessary to execute the init package function to register the "postgress" driver.
	// That makes the "postgress" driver available in the frist sql.Open call.
	_ "github.com/lib/pq"
)

// PQHandler is the handler to use with Postgress databases, it is an object that encapsulates the generic SQLHandler using "composition".
type PQHandler struct {
	*SQLHandler
}

// NewPQHandler is the constructor to build a handler for the Postgress dbtype
func NewPQHandler(connection string) (*PQHandler, error) {
	db, err := sql.Open("postgress", connection)
	return &PQHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
