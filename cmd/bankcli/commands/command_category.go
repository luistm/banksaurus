package commands

import (
	"fmt"

	"github.com/luistm/go-bank-cli/lib"

	"github.com/luistm/go-bank-cli/cmd/bankcli/configurations"
	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
	"github.com/luistm/go-bank-cli/lib/categories"
)

// Category command
type Category struct{}

// Execute the report command
func (c *Category) Execute(arguments map[string]interface{}) *Response {
	var out string
	var err error
	var cats []lib.Identifier

	dbName, dbPath := configurations.GetDatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return &Response{err: err, output: out}
	}
	defer SQLStorage.Close()

	categoriesInteractor := categories.NewInteractor(SQLStorage)

	if arguments["category"].(bool) && arguments["new"].(bool) {
		cats, err = categoriesInteractor.Create(arguments["<name>"].(string))
	}
	if arguments["category"].(bool) && arguments["show"].(bool) {
		cats, err = categoriesInteractor.GetAll()
	}

	for _, c := range cats {
		out += fmt.Sprintf("%s\n", c)
	}

	return &Response{err: err, output: out}
}
