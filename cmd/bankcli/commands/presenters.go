package commands

import (
	"fmt"

	"github.com/luistm/go-bank-cli/lib"
)

// CLIPresenter shows data in the command line
type CLIPresenter struct{}

// Present receives the data to be shown
func (c *CLIPresenter) Present(identifiers []lib.Entity) error {
	var out string
	for _, s := range identifiers {
		out += fmt.Sprintf("%s\n", s)
	}

	fmt.Print(out)
	return nil
}
