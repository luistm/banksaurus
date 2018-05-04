package commands

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/luistm/banksaurus/lib/transaction"

	"github.com/luistm/banksaurus/bank"
	"github.com/luistm/banksaurus/lib"
)

var errOutputPipeUndefined = errors.New("Output pipe is undefined")

// NewPresenter creates a new presenter object
func NewPresenter(output io.Writer) bank.Presenter {
	return &CLIPresenter{output: output}
}

// CLIPresenter shows data in the command line
type CLIPresenter struct {
	output io.Writer
}

// Present receives the data to be shown
func (c *CLIPresenter) Present(entities ...lib.Entity) error {

	if c.output == nil {
		return errOutputPipeUndefined
	}

	const padding = 1
	w := tabwriter.NewWriter(c.output, 0, 0, padding, ' ', 0)

	for _, entity := range entities {
		switch entity.(type) {
		case *transaction.Transaction:
			stringArray := strings.Fields(entity.String())
			value := stringArray[0]
			seller := strings.Join(stringArray[1:], " ")
			fmt.Fprintf(w, "%s\t%s\n", value, seller)
		default:
			fmt.Fprintf(w, "%s\n", entity)
		}
	}
	w.Flush()

	return nil
}
