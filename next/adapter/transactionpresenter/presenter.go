package transactionpresenter

import (
	"strconv"
)

// NewPresenter creates a new presenter instance
func NewPresenter() (*Presenter, error) {
	return &Presenter{}, nil
}

// Presenter repackages transactions into a view model
type Presenter struct {
	viewModel  *ViewModel
	OutputData []string
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

	p.OutputData = outputData

	return nil
}


// TODO: outputData should be read only. Make private and create a method to access it


// ViewModel returns an object with the data to be presented
func (p *Presenter) ViewModel() (*ViewModel, error) {
	// TODO: Handle error. Data shouldbe available only if cleaned by present
	p.viewModel, _ = NewViewModel(p.OutputData)
	return p.viewModel, nil
}
