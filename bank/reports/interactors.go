package reports

import (
	"github.com/luistm/go-bank-cli/bank/transactions"
	"github.com/luistm/go-bank-cli/infrastructure"
	"github.com/luistm/go-bank-cli/lib/customerrors"
)

// NewInteractor creates an interactor for reports
func NewInteractor(storage infrastructure.CSVStorage) *interactor {
	r := transactions.NewRepository(storage)
	return &interactor{repository: r}
}

type interactor struct {
	repository transactions.Repository
}

func (i *interactor) Report() (*Report, error) {
	if i.repository == nil {
		return &Report{}, customerrors.ErrRepositoryUndefined
	}

	r := &Report{}
	ts, err := i.repository.GetAll()
	if err != nil {
		return r, &customerrors.ErrRepository{Msg: err.Error()}
	}

	r.transactions = ts

	return r, nil
}
