package transactions

import (
	"github.com/luistm/go-bank-cli/bank"
	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
	"github.com/luistm/go-bank-cli/lib"
	"github.com/luistm/go-bank-cli/lib/customerrors"
	"github.com/luistm/go-bank-cli/lib/sellers"
)

// NewInteractor creates a new transactions interactor
func NewInteractor(s bank.CSVHandler) *Interactor {

	var DatabaseName = "bank.db"
	var DatabasePath = "/tmp"
	db, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		// TODO: Fix me
		panic(err)
	}

	sellersInteractor := sellers.NewInteractor(db)

	return &Interactor{
		repository: &repository{
			storage: s,
		},
		sellerInteractor: sellersInteractor,
	}
}

type Interactor struct {
	repository         iRepository
	sellerInteractor   lib.Creator
	categoryInteractor lib.Creator
}

// Load fetches raw data from a repository and processes it into objects
// to be persisted in storage.
func (i *Interactor) Load() error {

	if i.repository == nil {
		return customerrors.ErrRepositoryUndefined
	}

	transactions, err := i.repository.GetAll()
	if err != nil {
		return &customerrors.ErrRepository{Msg: err.Error()}
	}

	if i.sellerInteractor == nil {
		return customerrors.ErrInteractorUndefined
	}

	for _, t := range transactions {
		_, err = i.sellerInteractor.Create(t.s.String())
		if err != nil {
			return &customerrors.ErrInteractor{Msg: err.Error()}
		}
	}

	// if i.categoryInteractor == nil {
	// 	return customerrors.ErrInteractorUndefined
	// }

	// for _, t := range transactions {
	// 	_, err = i.categoryInteractor.Create(t.c.String())
	// 	if err != nil {
	// 		return &customerrors.ErrInteractor{Msg: err.Error()}
	// 	}
	// }

	return nil
}
