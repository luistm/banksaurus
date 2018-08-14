package listtransactions

import "github.com/luistm/banksaurus/next/entity/transaction"

// Presenter interface for formatting data for presentation
type Presenter interface {
	Present([]map[string]int64) error
}

// TransactionGateway is a collection of entities to be used by the listtransactions
type TransactionGateway interface {
	GetAll() ([]*transaction.Entity, error)
}
