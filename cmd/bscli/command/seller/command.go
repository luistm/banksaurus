package seller

import (
	"github.com/luistm/banksaurus/banksaurus/changesellername"
	"github.com/luistm/banksaurus/banksaurus/listsellers"
	"github.com/luistm/banksaurus/cmd/bscli/adapter/presenterlistsellers"
	"github.com/luistm/banksaurus/cmd/bscli/adapter/sellergateway"
	"github.com/luistm/banksaurus/cmd/bscli/application"
	"os"
)

// Command seller
type Command struct{}

// Execute the seller command with arguments
func (s *Command) Execute(arguments map[string]interface{}) error {

	db, err := application.Database()
	if err != nil {
		return err
	}
	defer db.Close()

	sr, err := sellergateway.NewSellerRepository(db)
	if err != nil {
		return err
	}

	if arguments["seller"].(bool) && arguments["show"].(bool) {

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
		i, err := changesellername.NewInteractor(sr)
		if err != nil {
			return err
		}

		r, err := changesellername.NewRequest(arguments["<id>"].(string), arguments["<name>"].(string))
		if err != nil {
			return err
		}

		err = i.Execute(r)
		if err != nil {
			return err
		}
	}

	return nil
}
