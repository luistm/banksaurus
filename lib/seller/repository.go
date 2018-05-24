package seller

import (
	"errors"
	"fmt"

	"github.com/luistm/banksaurus/lib"
	"github.com/mattn/go-sqlite3"
)

// NewRepository creates a Sellers ofr seller
func NewRepository(db lib.SQLInfrastructer) lib.Repository {
	return &Sellers{SQLStorage: db}
}

const (
	saveStatement   = "INSERT INTO seller(slug, name ) VALUES (?, ?)"
	updateStatement = "UPDATE seller SET name=? WHERE slug=?"
)

// Sellers repository
type Sellers struct {
	SQLStorage lib.SQLInfrastructer
}

// Save a seller
func (sr *Sellers) Save(ent lib.Entity) error {

	if sr.SQLStorage == nil {
		return lib.ErrInfrastructureUndefined
	}

	s := ent.(*Seller)
	err := sr.SQLStorage.Execute(saveStatement, s.slug, s.name)
	if err != nil {
		switch err.(type) {
		case sqlite3.Error:
			if err.(sqlite3.Error).Code == sqlite3.ErrConstraint {
				if err = sr.SQLStorage.Execute(updateStatement, s.name, s.slug); err != nil {
					return &lib.ErrInfrastructure{Msg: err.Error()}
				}
			}
		default:
			return &lib.ErrInfrastructure{Msg: err.Error()}
		}
	}

	return nil
}

// Get a seller
func (sr *Sellers) Get(sellerSlug string) (lib.Entity, error) {
	// statement := "SELECT * FROM seller WHERE slug=?"
	// rows, err := sr.SQLStorage.Query(statement, sellerSlug)
	// if err != nil {
	// 	return &Seller{}, fmt.Errorf("Database failure: %s", err)
	// }

	// var slug string
	// var name string
	// for rows.Next() {
	// 	err := rows.Scan(&slug, &name)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	break
	// }

	// return &Seller{slug: slug, name: name}, nil
	return &Seller{}, errors.New("get not implemented")
}

// GetAll fetches all sellers
func (sr *Sellers) GetAll() ([]lib.Entity, error) {
	statement := "SELECT * FROM seller"
	rows, err := sr.SQLStorage.Query(statement)
	if err != nil {
		return []lib.Entity{}, fmt.Errorf("database failure: %s", err)
	}
	defer rows.Close()

	var sellers []lib.Entity

	for rows.Next() {
		var slug string
		var name string
		err := rows.Scan(&slug, &name)
		if err != nil {
			return nil, err
		}
		sellers = append(sellers, &Seller{slug: slug, name: name})
	}

	return sellers, nil
}
