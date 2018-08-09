package transactionpresenter

import (
	"errors"
	"strconv"
)

// ErrNoDataToPresent ...
var ErrNoDataToPresent = errors.New("did not receive data to present")

// NewPresenter creates a new presenter instance
func NewPresenter() (*Presenter, error) {
	return &Presenter{}, nil
}

// Presenter repackages transactions into a view model
type Presenter struct {
	viewModel  *ViewModel
	outputData []string
	hasData    bool
}

// Present receives data from the interactor
// and transforms it in data suitable to be sent
// to the view model.
func (p *Presenter) Present(data []map[string]int64) error {
	outputData := []string{}
	for _, dict := range data {
		for key, value := range dict {
			valueInt := strconv.FormatInt(value, 10)
			outputData = append(outputData, key, valueInt)
		}
	}

	p.outputData = outputData
	p.hasData = true

	return nil
}

// OutputData returns the output data already converted to strings
func (p *Presenter) OutputData() ([]string, error) {
	if !p.hasData {
		return []string{}, ErrNoDataToPresent
	}
	return p.outputData, nil
}

// ViewModel returns an object with the data to be presented
func (p *Presenter) ViewModel() (*ViewModel, error) {
	// TODO: Handle error. Data shouldbe available only if cleaned by present
	p.viewModel, _ = NewViewModel(p.outputData)
	return p.viewModel, nil
}
