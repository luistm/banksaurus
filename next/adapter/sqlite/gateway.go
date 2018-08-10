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
	return &SellerRepository{db}, nil
}

// SellerRepository handles persistence in a database
type SellerRepository struct {
	db *sql.DB
}

// Saves seller to the database
func (sr *SellerRepository) Save(seller *seller.Entity) error {

	insertStatement := "INSERT INTO seller(slug) VALUES (?)"
	_, err := sr.db.Exec(insertStatement, seller.ID())
	if err != nil {
		return err
	}

	return nil
}
