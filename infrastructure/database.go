package infrastructure

import "database/sql"

// DatabaseHandler handles database operations
type DatabaseHandler struct {
	path string
	*sql.DB
}
