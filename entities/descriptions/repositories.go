package descriptions

import (
	"github.com/luistm/go-bank-cli/entities"
)

var saveStatement = "INSERT INTO descriptions(slug, friendlyName ) VALUES (?, ?)"

type repository struct {
	SQLStorage entities.SQLDatabaseHandler
}

func (r *repository) Save(d *Description) error {

	if r.SQLStorage == nil {
		return entities.ErrInfrastructureUndefined
	}

	err := r.SQLStorage.Execute(saveStatement, d.rawName, d.friendlyName)
	if err != nil {
		return &entities.ErrInfrastructure{Msg: err.Error()}
	}

	return nil
}
