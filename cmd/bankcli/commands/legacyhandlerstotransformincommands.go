package commands

import (
	"fmt"

	"github.com/luistm/go-bank-cli/cmd/bankcli/configurations"
	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
	"github.com/luistm/go-bank-cli/lib/categories"
	"github.com/luistm/go-bank-cli/lib/sellers"
)

// CreateCategoryHandler handles category creation command
func CreateCategoryHandler(name string) (string, error) {

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

func ShowCategoriesHandler() (string, error) {

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
