package commands

import (
	"errors"
	"fmt"
	"io"
	"text/tabwriter"

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
func (c *CLIPresenter) Present(identifiers ...lib.Entity) error {

	if c.output == nil {
		return errOutputPipeUndefined
	}

	const padding = 2
	w := tabwriter.NewWriter(c.output, 0, 0, padding, ' ', tabwriter.Debug)

	for _, s := range identifiers {
		fmt.Fprintf(w, "%s\n", s)
	}
	w.Flush()

	return nil
}
