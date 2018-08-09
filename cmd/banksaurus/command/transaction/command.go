package transaction

import (
	"os"

	"github.com/luistm/banksaurus/app"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
	"github.com/luistm/banksaurus/services/transaction"
)

// Command handles transaction command
type Command struct{}

// Executes the command instance
func (tc *Command) Execute(arguments map[string]interface{}) error {

	dbName, dbPath := app.DatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return err
	}
	defer SQLStorage.Close()

	// TODO: Get dependencies from the app here
	// SQLStorage := app.Get(app.config.storage)

	i, err := transaction.New(SQLStorage, NewPresenter(os.Stdout))
	if err != nil {
		return err
	}
	if err := i.Execute(); err != nil {
		return err
	}

	return err
}
