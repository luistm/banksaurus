package commands

import (
	"os"

	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
	"github.com/luistm/banksaurus/lib/categories"
)

// Category command
type Category struct{}

// Execute the report command
func (c *Category) Execute(arguments map[string]interface{}) error {
	var err error

	dbName, dbPath := configurations.DatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return err
	}
	defer SQLStorage.Close()

	categoriesInteractor := categories.NewInteractor(SQLStorage, NewPresenter(os.Stdout))

	if arguments["category"].(bool) && arguments["new"].(bool) {
		err = categoriesInteractor.Create(arguments["<name>"].(string))
	}
	if arguments["category"].(bool) && arguments["show"].(bool) {
		err = categoriesInteractor.GetAll()
	}

	if err != nil {
		return nil
	}

	return nil
}
