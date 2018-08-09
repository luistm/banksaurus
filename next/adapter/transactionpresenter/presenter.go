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
			valueString := strconv.FormatInt(value, 10)

			// Wooow... there must be a better way of doing this
			comma := rune(44)
			finalString := []rune{}
			a := []rune(valueString)
			for i := 0; i < len(a); i++ {
				if i == len(a)-2 {
					finalString = append(finalString, comma)
				}
				finalString = append(finalString, a[i])
			}

			outputData = append(outputData, string(finalString)+"€", key)
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
	if !p.hasData {
		return &ViewModel{}, ErrNoDataToPresent
	}

	viewModel, err := NewViewModel(p.outputData)
	if err != nil {
		return &ViewModel{}, err
	}

	return viewModel, nil
}
