package presenterlistsellers

import (
	"fmt"
	"io"
	"text/tabwriter"
)

func NewViewModel(data []string) (*ViewModel, error) {
	return &ViewModel{data}, nil
}

type ViewModel struct {
	data []string
}

func (vm *ViewModel) Write(view io.Writer) error {

	const padding = 1
	w := tabwriter.NewWriter(view, 0, 0, padding, ' ', 0)

	for _, seller := range vm.data {
		fmt.Fprintf(w, "%s\n", seller)
	}

	w.Flush()

	return nil
}
