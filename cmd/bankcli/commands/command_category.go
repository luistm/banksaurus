package commands

import (
	"fmt"

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

	if arguments["category"].(bool) && arguments["new"].(bool) {
		out, err = c.createCategoryHandler(arguments["<name>"].(string))
	}
	if arguments["category"].(bool) && arguments["show"].(bool) {
		out, err = c.showCategoriesHandler()
	}

	return &Response{err: err, output: out}
}

func (c *Category) createCategoryHandler(name string) (string, error) {

	dbName, dbPath := configurations.GetDatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return "", err
	}
	defer SQLStorage.Close()

	categoriesInteractor := categories.NewInteractor(SQLStorage)
	cats, err := categoriesInteractor.Create(name)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("Created category '%s'", cats[0])
	return msg, nil
}

func (c *Category) showCategoriesHandler() (string, error) {

	dbName, dbPath := configurations.GetDatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return "", err
	}
	defer SQLStorage.Close()

	categoriesInteractor := categories.NewInteractor(SQLStorage)
	cats, err := categoriesInteractor.GetAll()
	if err != nil {
		return "", err
	}

	out := ""
	for _, c := range cats {
		out += fmt.Sprintf("%s\n", c)
	}

	return out, nil
}
