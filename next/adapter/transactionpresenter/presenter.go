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
	viewModel *ViewModel
}

// Present receives data from the interactor
func (p *Presenter) Present(data []map[string]int64) error {
	rawData := []string{}
	for _, dict := range data {
		for key, value := range dict {
			valueInt := strconv.FormatInt(value, 10)
			rawData = append(rawData, key, valueInt)
		}
	}

	p.viewModel = &ViewModel{rawData}

	return nil
}

// ViewModel returns an object with the data to be presented
func (p *Presenter) ViewModel() (*ViewModel, error) {
	return p.viewModel, nil
}
