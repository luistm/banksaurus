package commands

import (
	"github.com/luistm/banksaurus/bank/transactions_show"
	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
)

// TransactionCommand handles transaction
type TransactionCommand struct{}

// Executes the command instance
func (tc *TransactionCommand) Execute(arguments map[string]interface{}) error {

	dbName, dbPath := configurations.DatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return err
	}
	defer SQLStorage.Close()

	i, err := transactions_show.New()
	if err != nil {
		return err
	}
	if err := i.Execute(); err != nil {
		return err
	}

	return err
}
