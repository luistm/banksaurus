package reports

import (
	"github.com/luistm/go-bank-cli/bank/transactions"
	"github.com/luistm/go-bank-cli/lib/customerrors"
)

// NewInteractor creates an Interactor for reports
func NewInteractor(r transactions.Fetcher) *Interactor {
	return &Interactor{repository: r}
}

type Interactor struct {
	repository transactions.Fetcher
}

func (i *Interactor) ReportFromRecords() (*Report, error) {
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
