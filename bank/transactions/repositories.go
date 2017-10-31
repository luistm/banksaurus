package transactions

import (
	"github.com/luistm/go-bank-cli/bank"
	"github.com/luistm/go-bank-cli/lib/customerrors"
)

type repository struct {
	storage bank.CSVHandler
}

func (r *repository) GetAll() ([]*Transaction, error) {

	if r.storage == nil {
		return []*Transaction{}, customerrors.ErrInfrastructureUndefined
	}

	_, err := r.storage.Lines()
	if err != nil {
		return []*Transaction{}, &customerrors.ErrInfrastructure{Msg: err.Error()}
	}

	return []*Transaction{}, nil
}
