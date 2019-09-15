package presenterlisttransactions

import (
	"errors"
	"github.com/luistm/banksaurus/banksauruslib/entities/transaction"
	"strings"
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
func (p *Presenter) Present(data []map[string]*transaction.Money) error {
	outputData := []string{}
	comma := rune(44)
	zero := rune(48)

	for _, dict := range data {
		for key, value := range dict {

			// TODO: create a test with for each possible combination of input
			valueString := value.String()

			// Wooow... there must be a better way of doing this
			finalString := []rune{}

			// Converting two digit values.
			if len(valueString) == 2 {
				// 12 becomes 0,12
				finalString = append(finalString, zero, comma)
				valueString = string(finalString) + valueString

			} else if len(valueString) == 3 && strings.HasPrefix(valueString, "-") {
				// -12 becomes -0,12
				finalString = append(finalString, zero, comma)
				prefix := "-"
				valueString = prefix + string(finalString) + strings.TrimPrefix(valueString, prefix)

			} else {
				a := []rune(valueString)
				for i := 0; i < len(a); i++ {
					if i == len(a)-2 {
						finalString = append(finalString, comma)
					}
					finalString = append(finalString, a[i])
				}
				valueString = string(finalString)
			}

			outputData = append(outputData, valueString+"€", key)
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
