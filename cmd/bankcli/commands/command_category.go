package commands

import (
	"os"

	"github.com/luistm/go-bank-cli/cmd/bankcli/configurations"
	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
	"github.com/luistm/go-bank-cli/lib/categories"
)

// Category command
type Category struct{}

// Execute the report command
func (c *Category) Execute(arguments map[string]interface{}) *Response {
	var err error

	dbName, dbPath := configurations.GetDatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return &Response{err: err}
	}
	defer SQLStorage.Close()

	categoriesInteractor := categories.NewInteractor(SQLStorage, NewPresenter(os.Stdout))

	if arguments["category"].(bool) && arguments["new"].(bool) {
		err = categoriesInteractor.Create(arguments["<name>"].(string))
	}
	if arguments["category"].(bool) && arguments["show"].(bool) {
		err = categoriesInteractor.GetAll()
	}

	return &Response{err: err}
}
