package sqlite

import (
	"database/sql"
	"errors"
	"github.com/luistm/banksaurus/next/entity/seller"
)

// ErrDatabaseUndefined ...
var ErrDatabaseUndefined = errors.New("database is not defined")

// NewSellerRepository creates a new seller repository instance
func NewSellerRepository(db *sql.DB) (*SellerRepository, error) {
	if db == nil {
		return &SellerRepository{}, ErrDatabaseUndefined
	}
	return &SellerRepository{}, nil
}

// SellerRepository handles persistence in a database
type SellerRepository struct{}

// Saves seller to the database
func (*SellerRepository) Save(entity *seller.Entity) error {
	panic("implement me")
}
