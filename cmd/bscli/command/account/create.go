package account

import (
	"github.com/luistm/banksaurus/banksaurus/createaccount"
	"github.com/luistm/banksaurus/cmd/bscli/adapter/accountgateway"
	"github.com/luistm/banksaurus/cmd/bscli/application"
)

// CreateAccountCommand for account
type CreateAccountCommand struct{}

// Execute the account command
func (c *CreateAccountCommand) Execute(input map[string]interface{}) error {

	db, err := application.Database()
	if err != nil {
		return err
	}

	repository, err := accountgateway.NewRepository(db)
	if err != nil {
		return err
	}

	i, err := createaccount.NewInteractor(repository)
	if err != nil {
		return err
	}

	request, err := createaccount.NewRequest()
	if err != nil {
		return err
	}

	err = i.Execute(request)
	if err != nil {
		return err
	}

	return nil
}
