package descriptions

import (
	"github.com/luistm/go-bank-cli/infrastructure"
)

type repository struct {
	DBHandler infrastructure.Storage
}

func (r *repository) Save(d *Description) error {

	return nil
}
