package presenterlisttransactions

import (
	"errors"
	"fmt"
	"text/tabwriter"
)

// ErrDataHasOddLength ...
var ErrDataHasOddLength = errors.New("data has odd length")

// NewViewModel creates a view model instance
func NewViewModel(data []string) (*ViewModel, error) {

	if len(data)%2 == 1 {
		return &ViewModel{}, ErrDataHasOddLength
	}

	return &ViewModel{raw: data}, nil
}

// ViewModel contains data to be shown.
type ViewModel struct {
	raw []string
}

// Writes the data into a view
func (vm *ViewModel) Write(view Viewer) {

	const padding = 1
	w := tabwriter.NewWriter(view, 0, 0, padding, ' ', 0)

	for i := 0; i < len(vm.raw); i = i + 2 {
		fmt.Fprintf(w, "%s\t%s\n", vm.raw[i], vm.raw[i+1])
	}

	w.Flush()
}
