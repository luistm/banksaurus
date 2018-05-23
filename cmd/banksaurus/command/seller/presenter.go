package seller

import (
	"errors"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/luistm/banksaurus/lib"
	"github.com/luistm/banksaurus/services"
)

var errOutputPipeUndefined = errors.New("output pipe is undefined")

// NewPresenter creates a new presenter object
func NewPresenter(output io.Writer) services.Presenter {
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
		fmt.Fprintf(w, "%s\n", entity)
	}
	w.Flush()

	return nil
}
