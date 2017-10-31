package transactions

import (
	"github.com/luistm/go-bank-cli/infrastructure"
	"github.com/luistm/go-bank-cli/lib/customerrors"
)

type repository struct {
	storage infrastructure.CSVStorage
}

func (r *repository) GetAll() ([]*Transaction, error) {

	if r.storage == nil {
		return []*Transaction{}, customerrors.ErrInfrastructureUndefined
	}

	return []*Transaction{}, nil
}
