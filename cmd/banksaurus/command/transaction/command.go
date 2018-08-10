package transaction

import (
	"github.com/luistm/banksaurus/app"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
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

	return err
}
