package reportfromrecords

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/seller"
	"github.com/luistm/banksaurus/lib/transaction"
)

// New creates a new ReportFromRecords use case
func New(
	transactionsRepository lib.Repository, sellersRepository lib.Repository, presenter lib.Presenter,
) (*ReportFromRecords, error) {

	if transactionsRepository == nil ||
		sellersRepository == nil {
		return &ReportFromRecords{}, lib.ErrRepositoryUndefined
	}
	if presenter == nil {
		return &ReportFromRecords{}, lib.ErrPresenterUndefined
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

	var ts []lib.Entity

	transactionsList, err := i.transactionsRepository.GetAll()
	if err != nil {
		return &lib.ErrRepository{Msg: err.Error()}
	}
	if len(transactionsList) == 0 {
		return nil
	}

	for _, t := range transactionsList {
		// FIXME: For each transaction, fetch only the needed seller, not all the seller
		allSellers, err := i.sellersRepository.GetAll()
		if err != nil {
			return &lib.ErrRepository{Msg: err.Error()}
		}

		for _, s := range allSellers {
			if s.ID() == t.(*transaction.Transaction).Seller.ID() {
				t.(*transaction.Transaction).Seller = s.(*seller.Seller)
				break
			}
		}
		ts = append(ts, t.(*transaction.Transaction))
	}

	if err := i.presenter.Present(ts...); err != nil {
		return &lib.ErrPresenter{Msg: err.Error()}
	}

	return nil
}
