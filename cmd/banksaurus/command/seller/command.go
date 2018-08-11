package seller

import (
	"github.com/luistm/banksaurus/next/application/infrastructure/relational"

	"github.com/luistm/banksaurus/next/application/adapter/presenterlistsellers"
	"github.com/luistm/banksaurus/next/application/adapter/sqlite"
	"github.com/luistm/banksaurus/next/listsellers"
	"os"
)

// Command seller
type Command struct{}

// Execute the seller command with arguments
func (s *Command) Execute(arguments map[string]interface{}) error {

	if arguments["seller"].(bool) && arguments["new"].(bool) {
		panic("seller new not implemented")
	}

	if arguments["seller"].(bool) && arguments["show"].(bool) {

		db, err := relational.NewDatabase()
		if err != nil {
			return err
		}

		sr, err := sqlite.NewSellerRepository(db)
		if err != nil {
			return err
		}

		p, err := presenterlistsellers.NewPresenter()
		if err != nil {
			return err
		}

		i, err := listsellers.NewInteractor(sr, p)
		if err != nil {
			return err
		}

		err = i.Execute()
		if err != nil {
			return err
		}

		vm, err := p.ViewModel()
		if err != nil {
			return err
		}

		vm.Write(os.Stdout)
	}

	if arguments["seller"].(bool) && arguments["change"].(bool) {
		panic("seller change not implemented")
	}

	return nil
}
