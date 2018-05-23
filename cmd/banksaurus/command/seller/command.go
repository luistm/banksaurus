package seller

import (
	"os"

	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/infrastructure/sqlite"
	"github.com/luistm/banksaurus/services/seller"
)

// Command command
type Command struct{}

// Execute the seller command with arguments
func (s *Command) Execute(arguments map[string]interface{}) error {
	var err error

	dbName, dbPath := configurations.DatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return err
	}
	defer SQLStorage.Close()

	sellersInteractor := seller.New(SQLStorage, NewPresenter(os.Stdout))

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
