package reportfromrecordsgrouped

import "github.com/luistm/banksaurus/lib/customerrors"
import "github.com/luistm/banksaurus/lib"

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