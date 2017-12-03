package sellers

import (
	"fmt"

	"github.com/mattn/go-sqlite3"

	"github.com/luistm/go-bank-cli/lib"
	"github.com/luistm/go-bank-cli/lib/customerrors"
)

var saveStatement = "INSERT INTO sellers(slug, name ) VALUES (?, ?)"
var updateStatement = "UPDATE sellers SET name=? WHERE slug=?"

// NewRepository creates a repository ofr sellers
func NewRepository(db lib.SQLInfrastructer) lib.Repository {
	return &repository{SQLStorage: db}
}

type repository struct {
	SQLStorage lib.SQLInfrastructer
}

func (r *repository) Save(ent lib.Identifier) error {

	if r.SQLStorage == nil {
		return customerrors.ErrInfrastructureUndefined
	}

	s := ent.(*Seller)
	err := r.SQLStorage.Execute(saveStatement, s.slug, s.name)
	if err != nil {
		switch err.(type) {
		case sqlite3.Error:
			if err.(sqlite3.Error).Code == sqlite3.ErrConstraint {
				if err = r.SQLStorage.Execute(updateStatement, s.name, s.slug); err != nil {
					return &customerrors.ErrInfrastructure{Msg: err.Error()}
				}
			}
		default:
			return &customerrors.ErrInfrastructure{Msg: err.Error()}
		}
	}

	return nil
}

func (r *repository) Get(sellerSlug string) (lib.Identifier, error) {
	// statement := "SELECT * FROM sellers WHERE slug=?"
	// rows, err := r.SQLStorage.Query(statement, sellerSlug)
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
	return &Seller{}, nil
}

// GetAll fetches all sellers
func (r *repository) GetAll() ([]lib.Identifier, error) {
	statement := "SELECT * FROM sellers"
	rows, err := r.SQLStorage.Query(statement)
	if err != nil {
		return []lib.Identifier{}, fmt.Errorf("Database failure: %s", err)
	}
	defer rows.Close()

	sellers := []lib.Identifier{}

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
