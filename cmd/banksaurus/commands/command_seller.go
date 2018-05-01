package commands

import (
	"os"

	"github.com/luistm/banksaurus/bank"
	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
)

// Seller commands
type Seller struct{}

// Execute the seller command with arguments
func (s *Seller) Execute(arguments map[string]interface{}) error {
	var err error

	dbName, dbPath := configurations.DatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return err
	}
	defer SQLStorage.Close()

	sellersInteractor := bank.NewInteractor(SQLStorage, NewPresenter(os.Stdout))

	if arguments["seller"].(bool) && arguments["new"].(bool) {
		err = sellersInteractor.Create(arguments["<name>"].(string))
	}

	if arguments["seller"].(bool) && arguments["show"].(bool) {
		err = sellersInteractor.GetAll()
	}

	if arguments["seller"].(bool) && arguments["change"].(bool) {
		err = sellersInteractor.Update(arguments["<id>"].(string), arguments["<name>"].(string))
	}

	if err != nil {
		return nil
	}

	return nil
}
