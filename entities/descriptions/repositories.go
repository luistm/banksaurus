package descriptions

import (
	"github.com/luistm/go-bank-cli/entities"
)

var saveStatement = "INSERT INTO descriptions(slug, friendlyName ) VALUES (?, ?)"

type repository struct {
	SQLStorage entities.SQLDatabaseHandler
}

func (r *repository) Save(ent entities.Entity) error {

	if r.SQLStorage == nil {
		return entities.ErrInfrastructureUndefined
	}

	d := ent.(*Description)
	err := r.SQLStorage.Execute(saveStatement, d.rawName, d.friendlyName)
	if err != nil {
		return &entities.ErrInfrastructure{Msg: err.Error()}
	}

	return nil
}

func (r *repository) Get(d string) (entities.Entity, error) {
	return &Description{}, nil
}

func (r *repository) GetAll() ([]entities.Entity, error) {
	return []entities.Entity{}, nil
}
