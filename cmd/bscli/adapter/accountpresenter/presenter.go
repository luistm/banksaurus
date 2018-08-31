package accountpresenter

import (
	"errors"
	"github.com/luistm/banksaurus/money"
	"strconv"
)

var ErrNoDataAvailable = errors.New("no data to show")

func New() (*Presenter, error) {
	return &Presenter{}, nil
}

type Presenter struct {
	hasData bool
	m       *money.Money
}

func (p *Presenter) Present(m *money.Money) error {
	p.m = m
	p.hasData = true

	return nil
}

func (p *Presenter) ViewModel() (*ViewModel, error) {
	if !p.hasData {
		return &ViewModel{}, ErrNoDataAvailable
	}

	value := strconv.FormatInt(p.m.Value(), 10)
	vm, err := NewViewModel(value)
	if err != nil {
		return &ViewModel{}, err
	}

	return vm, nil
}
