package presenterlistsellers

import "errors"

// ErrNoDataToPresent ...
var ErrNoDataToPresent = errors.New("did not receive data to present")

// NewPresenter ...
func NewPresenter() (*Presenter, error) {
	return &Presenter{}, nil
}

// Presenter to list sellers
type Presenter struct {
	raw     []string
	hasData bool
}

func (p *Presenter) Present(sellers []string) error {
	p.raw = sellers
	p.hasData = true

	return nil
}

func (p *Presenter) OutputData() ([]string, error) {
	if !p.hasData {
		return []string{}, ErrNoDataToPresent
	}
	return p.raw, nil
}

func (p *Presenter) ViewModel() (*ViewModel, error) {
	if !p.hasData {
		return &ViewModel{}, ErrNoDataToPresent
	}

	vm, err := NewViewModel(p.raw)
	if err != nil {
		return &ViewModel{}, err
	}

	return vm, nil
}
