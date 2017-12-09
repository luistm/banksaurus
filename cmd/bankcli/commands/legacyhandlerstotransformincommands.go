package commands

import (
	"fmt"

	"github.com/luistm/go-bank-cli/infrastructure/sqlite"
	"github.com/luistm/go-bank-cli/lib/categories"
	"github.com/luistm/go-bank-cli/lib/sellers"
)

var DatabaseName = "bank.db"
var DatabasePath = "/tmp"

// createCategoryHandler handles category creation command
func CreateCategoryHandler(name string) (string, error) {

	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
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
	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
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
	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return out, err
	}
	defer SQLStorage.Close()

	sellersInteractor := sellers.NewInteractor(SQLStorage)
	s, err := sellersInteractor.Create(name)
	if err != nil {
		return out, err
	}

	return s.String(), nil
}

func ShowSellersHandler() (string, error) {
	var out string
	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return out, err
	}
	defer SQLStorage.Close()

	sellersInteractor := sellers.NewInteractor(SQLStorage)
	sellers, err := sellersInteractor.GetAll()
	if err != nil {
		return out, err
	}

	for _, s := range sellers {
		out += fmt.Sprintf("%s\n", s)
	}

	return out, nil
}

func SellerChangePrettyName(sellerID string, name string) (string, error) {
	var out string
	SQLStorage, err := sqlite.New(DatabasePath, DatabaseName, false)
	if err != nil {
		return out, err
	}
	defer SQLStorage.Close()

	sellersInteractor := sellers.NewInteractor(SQLStorage)
	err = sellersInteractor.Update(sellerID, name)
	if err != nil {
		return out, err
	}

	return out, nil
}
