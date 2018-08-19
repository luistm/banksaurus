package databasegateway

import (
	"database/sql"
	"errors"
)

// ErrDatabaseUndefined ...
var ErrDatabaseUndefined = errors.New("database is not defined")

// NewSellerRepository creates a new seller repository instance
func NewSellerRepository(db *sql.DB) (*Repository, error) {
	if db == nil {
		return &Repository{}, ErrDatabaseUndefined
	}
	return &Repository{db}, nil
}

// Repository handles persistence in a database
type Repository struct {
	db *sql.DB
}
