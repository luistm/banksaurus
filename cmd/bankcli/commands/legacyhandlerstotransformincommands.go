package commands

import (
	"github.com/luistm/go-bank-cli/cmd/bankcli/configurations"
	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
	"github.com/luistm/go-bank-cli/lib/sellers"
)

func CreateSellerHandler(name string) (string, error) {
	var out string
	dbName, dbPath := configurations.GetDatabasePath()
	SQLStorage, err := sqlite.New(dbPath, dbName, false)
	if err != nil {
		return out, err
	}
	defer SQLStorage.Close()

	sellersInteractor := sellers.NewInteractor(SQLStorage, nil)
	s, err := sellersInteractor.Create(name)
	if err != nil {
		return out, err
	}

	return s.String(), nil
}

func ShowSellersHandler() error {
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

func SellerChangePrettyName(sellerID string, name string) (string, error) {
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
