package infrastructure

import "database/sql"

// DatabaseHandler handles database operations
type DatabaseHandler struct {
	path string
	*sql.DB
}

// Execute an sql statement
func (dh *DatabaseHandler) Execute(statement string) error {
	return nil
}

// Query fetches data from the database
func (dh *DatabaseHandler) Query(statement string) (*sql.Rows, error) {
	return nil, nil
}
