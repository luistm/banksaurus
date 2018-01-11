package commands

import (
	"fmt"

	"github.com/luistm/go-bank-cli/lib"
)

// CLIPresenter shows data in the command line
type CLIPresenter struct{}

// Present receives the data to be shown
func (c *CLIPresenter) Present(identifiers ...lib.Entity) error {
	for _, s := range identifiers {
		fmt.Println(s.String())
	}

	return nil
}
