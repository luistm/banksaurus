package transaction

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/lib/transaction"
	"github.com/luistm/banksaurus/services"
)

// New crates a service instance
func New(storage lib.SQLInfrastructer, presenter services.Presenter) (services.Servicer, error) {
	r := transaction.NewRepository(nil, storage)
	return &Service{presenter, r}, nil
}

// Service shows transactions
type Service struct {
	presenter    services.Presenter
	transactions lib.Fetcher
}

// Execute the service
func (ts *Service) Execute() error {

	return nil
}
