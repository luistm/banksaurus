package sellers

import (
	"fmt"

	"github.com/luistm/go-bank-cli/lib"
	"github.com/luistm/go-bank-cli/lib/customerrors"
)

var saveStatement = "INSERT INTO sellers(slug, name ) VALUES (?, ?)"

type repository struct {
	SQLStorage lib.SQLDatabaseHandler
}

func (r *repository) Save(ent lib.Entity) error {

	if r.SQLStorage == nil {
		return customerrors.ErrInfrastructureUndefined
	}

	s := ent.(*Seller)
	err := r.SQLStorage.Execute(saveStatement, s.slug, s.name)
	if err != nil {
		return &customerrors.ErrInfrastructure{Msg: err.Error()}
	}

	return nil
}

func (r *repository) Get(s string) (lib.Entity, error) {
	return &Seller{}, nil
}

// GetAll fetches all sellers
func (r *repository) GetAll() ([]lib.Entity, error) {
	statement := "SELECT * FROM sellers"
	rows, err := r.SQLStorage.Query(statement)
	if err != nil {
		return []lib.Entity{}, fmt.Errorf("Database failure: %s", err)
	}

	sellers := []lib.Entity{}

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
