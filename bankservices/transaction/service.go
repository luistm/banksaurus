package transaction

import "github.com/luistm/banksaurus/bankservices"

// New creates a transaction show interactor
func New() (bankservices.Servicer, error) {
	return &Service{}, nil
}

// Service shows transaction
type Service struct{}

// Execute ...
func (ts *Service) Execute() error {
	return nil
}
