package report

import "github.com/luistm/banksaurus/next/entity/transaction"

// Presenter interface for formatting data for presentation
type Presenter interface {
	Present([]map[string]int64) error
}

// Repository is a collection of entities to be used by the report
type Repository interface {
	GetAll() ([]*transaction.Entity, error)
}
