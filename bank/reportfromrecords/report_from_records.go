package reportfromrecords

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/customerrors"
	"github.com/luistm/banksaurus/lib/sellers"
	"github.com/luistm/banksaurus/lib/transactions"
)

// New creates a new ReportFromRecords use case
func New(
	transactionsRepository lib.Repository, sellersRepository lib.Repository, presenter lib.Presenter,
) (*ReportFromRecords, error) {

	if transactionsRepository == nil ||
		sellersRepository == nil {
		return &ReportFromRecords{}, customerrors.ErrRepositoryUndefined
	}
	if presenter == nil {
		return &ReportFromRecords{}, customerrors.ErrPresenterUndefined
	}

	return &ReportFromRecords{
		transactionsRepository: transactionsRepository,
		sellersRepository:      sellersRepository,
		presenter:              presenter,
	}, nil
}

// ReportFromRecords makes a report from an input file.
// If a Seller has a pretty name, that name will be used.
type ReportFromRecords struct {
	transactionsRepository lib.Repository
	sellersRepository      lib.Repository
	presenter              lib.Presenter
}

// Execute ...
func (i *ReportFromRecords) Execute() error {

	ts := []lib.Entity{}

	transactionsList, err := i.transactionsRepository.GetAll()
	if err != nil {
		return &customerrors.ErrRepository{Msg: err.Error()}
	}

	for _, t := range transactionsList {
		// FIXME: For each transaction, fetch only the needed sellers, not all the sellers
		allSellers, err := i.sellersRepository.GetAll()
		if err != nil {
			return &customerrors.ErrRepository{Msg: err.Error()}
		}

		for _, s := range allSellers {
			if s.ID() == t.(*transactions.Transaction).Seller.ID() {
				t.(*transactions.Transaction).Seller = s.(*sellers.Seller)
				break
			}
		}
		ts = append(ts, t.(*transactions.Transaction))
	}

	if err := i.presenter.Present(ts...); err != nil {
		return &customerrors.ErrPresenter{Msg: err.Error()}
	}

	return nil
}
