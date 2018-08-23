package accountgateway

import (
	"database/sql"
	"github.com/luistm/banksaurus/account"
	"github.com/luistm/banksaurus/money"
	"github.com/pkg/errors"
)

// ErrDatabaseUndefined ...
var ErrDatabaseUndefined = errors.New("database is undefined")

// NewRepository ...
func NewRepository(db *sql.DB) (*Repository, error) {
	if db == nil {
		return &Repository{db}, ErrDatabaseUndefined
	}
	return &Repository{db}, nil
}

// Repository for account
type Repository struct {
	db *sql.DB
}

// GetByID returns the account with matches the account id
func (r *Repository) GetByID(id string) (*account.Entity, error) {
	return &account.Entity{}, nil
}

// New creates a new account entity
func (r *Repository) New(*money.Money) (*account.Entity, error) {
	return &account.Entity{}, nil
}
