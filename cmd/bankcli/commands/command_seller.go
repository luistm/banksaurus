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

	if arguments["seller"].(bool) && arguments["new"].(bool) {
		out, err = s.createSellerHandler(arguments["<name>"].(string))
	}

	if arguments["seller"].(bool) && arguments["show"].(bool) {
		err = s.showSellersHandler()
	}

	if arguments["seller"].(bool) && arguments["change"].(bool) {
		out, err = s.sellerChangePrettyName(
			arguments["<id>"].(string),
			arguments["<name>"].(string),
		)
	}

	return &Response{err: err, output: out}
}

func (s *Seller) createSellerHandler(name string) (string, error) {
	var out string

	dbName, dbPath := configurations.GetDatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return out, err
	}
	defer SQLStorage.Close()

	sellersInteractor := sellers.NewInteractor(SQLStorage, nil)
	_, err = sellersInteractor.Create(name)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (s *Seller) showSellersHandler() error {
	dbName, dbPath := configurations.GetDatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return err
	}
	defer SQLStorage.Close()

	presenter := &CLIPresenter{}

	sellersInteractor := sellers.NewInteractor(SQLStorage, presenter)
	err = sellersInteractor.GetAll()
	if err != nil {
		return err
	}

	return nil
}

func (s *Seller) sellerChangePrettyName(sellerID string, name string) (string, error) {
	var out string
	dbName, dbPath := configurations.GetDatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return out, err
	}
	defer SQLStorage.Close()

	sellersInteractor := sellers.NewInteractor(SQLStorage, nil)
	err = sellersInteractor.Update(sellerID, name)
	if err != nil {
		return out, err
	}

	return out, nil
}
