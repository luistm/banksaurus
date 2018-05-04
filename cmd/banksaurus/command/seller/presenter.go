package seller

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/transaction"
	"github.com/luistm/banksaurus/bankservices"
)

var errOutputPipeUndefined = errors.New("output pipe is undefined")

// NewPresenter creates a new presenter object
func NewPresenter(output io.Writer) bankservices.Presenter {
	return &CLIPresenter{output: output}
}

// CLIPresenter shows data in the command line
type CLIPresenter struct {
	output io.Writer
}

// Present receives the data to be shown
func (c *CLIPresenter) Present(entities ...banklib.Entity) error {

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
