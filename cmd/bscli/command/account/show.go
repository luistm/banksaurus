package account

import (
	"github.com/luistm/banksaurus/banksaurus/showaccount"
	"github.com/luistm/banksaurus/cmd/bscli/adapter/accountgateway"
	"github.com/luistm/banksaurus/cmd/bscli/application"
)

// ListAccountsCommand for account
type ListAccountsCommand struct{}

// Execute the account command
func (*ListAccountsCommand) Execute(input map[string]interface{}) error {
	db, err := application.Database()
	if err != nil {
		return err
	}

	repository, err := accountgateway.NewRepository(db)
	if err != nil {
		return err
	}

	i, err := showaccount.NewInteractor(repository, presenter)
	if err != nil {
		return err
	}

	request, err := showaccount.NewRequest(input["<name>"].(string))
	if err != nil {
		return err
	}

	err = i.Execute(request)
	if err != nil {
		return err
	}

	return nil
}
