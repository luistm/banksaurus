package commands

import (
	"github.com/luistm/go-bank-cli/cmd/bankcli/configurations"
	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
	"github.com/luistm/go-bank-cli/lib/sellers"
)

// Seller commands
type Seller struct{}

// Execute the seller command with arguments
func (s *Seller) Execute(arguments map[string]interface{}) *Response {
	var out string
	var err error

	dbName, dbPath := configurations.GetDatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return &Response{err: err, output: out}
	}
	defer SQLStorage.Close()
	presenter := &CLIPresenter{}

	sellersInteractor := sellers.NewInteractor(SQLStorage, presenter)

	if arguments["seller"].(bool) && arguments["new"].(bool) {
		_, err = sellersInteractor.Create(arguments["<name>"].(string))
	}

	if arguments["seller"].(bool) && arguments["show"].(bool) {
		err = sellersInteractor.GetAll()
	}

	if arguments["seller"].(bool) && arguments["change"].(bool) {
		err = sellersInteractor.Update(arguments["<id>"].(string), arguments["<name>"].(string))
	}

	return &Response{err: err, output: out}
}
