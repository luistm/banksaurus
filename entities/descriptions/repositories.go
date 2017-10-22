package descriptions

import (
	"github.com/luistm/go-bank-cli/infrastructure"
)

type repository struct {
	SQLStorage infrastructure.SQLStorage
}

func (r *repository) Save(d *Description) error {

	return nil
}
