package commands

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/luistm/go-bank-cli/bank/transactions"

	"github.com/luistm/go-bank-cli/lib"
)

var errOutputPipeUndefined = errors.New("Output pipe is undefined")

// NewPresenter creates a new presenter object
func NewPresenter(output io.Writer) lib.Presenter {
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
		case *transactions.Transaction:
			stringArray := strings.Fields(entity.String())
			price := stringArray[0]
			fmt.Fprintf(w, "%s\t%s\n", price, strings.Join(stringArray[1:], " "))
		default:
			fmt.Fprintf(w, "%s\n", entity)
		}
	}
	w.Flush()

	return nil
}
