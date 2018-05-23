package transaction

import (
	"os"

	"github.com/luistm/banksaurus/services/transaction"
	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
)

// Command handles transaction command
type Command struct{}

// Executes the command instance
func (tc *Command) Execute(arguments map[string]interface{}) error {

	dbName, dbPath := configurations.DatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return err
	}
	defer SQLStorage.Close()

	i, err := transaction.New(SQLStorage, NewPresenter(os.Stdout))
	if err != nil {
		return err
	}
	if err := i.Execute(); err != nil {
		return err
	}

	return err
}
