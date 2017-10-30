package sellers

import (
	"fmt"

	"github.com/luistm/go-bank-cli/entities"
)

var saveStatement = "INSERT INTO sellers(slug, name ) VALUES (?, ?)"

type repository struct {
	SQLStorage entities.SQLDatabaseHandler
}

func (r *repository) Save(ent entities.Entity) error {

	if r.SQLStorage == nil {
		return entities.ErrInfrastructureUndefined
	}

	s := ent.(*Seller)
	err := r.SQLStorage.Execute(saveStatement, s.slug, s.name)
	if err != nil {
		return &entities.ErrInfrastructure{Msg: err.Error()}
	}

	return nil
}

func (r *repository) Get(s string) (entities.Entity, error) {
	return &Seller{}, nil
}

// GetAll fetches all sellers
func (r *repository) GetAll() ([]entities.Entity, error) {
	statement := "SELECT * FROM sellers"
	rows, err := r.SQLStorage.Query(statement)
	if err != nil {
		return []entities.Entity{}, fmt.Errorf("Database failure: %s", err)
	}

	sellers := []entities.Entity{}

	for rows.Next() {
		var slug int
		var name string
		err := rows.Scan(&slug, &name)
		if err != nil {
			return nil, err
		}
		sellers = append(sellers, &Seller{})
	}

	return sellers, nil
}
