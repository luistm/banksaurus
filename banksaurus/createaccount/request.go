package createaccount

import (
	"errors"
	"github.com/luistm/banksaurus/money"
)

var ErrInvalidData = errors.New("invalid data")

func NewRequest(balance int64) (*Request, error) {
	m, err := money.NewMoney(balance)
	if err != nil {
		return &Request{}, err
	}

	return &Request{m, true}, nil
}

type Request struct {
	m        *money.Money
	hasMoney bool
}

func (r *Request) Balance() (*money.Money, error) {
	if !r.hasMoney {
		return &money.Money{}, ErrInvalidData
	}
	return r.m, nil
}
