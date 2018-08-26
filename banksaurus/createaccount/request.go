package createaccount

import (
	"errors"
	"github.com/luistm/banksaurus/money"
	"strconv"
	"strings"
)

var ErrInvalidData = errors.New("invalid data")

func NewRequest(input string) (*Request, error) {
	if input == "" {
		return &Request{}, ErrInvalidData
	}

	parsedInput := strings.Replace(input, ".", "", -1)
	parsedInput = strings.Replace(parsedInput, ",", "", -1)

	balance, err := strconv.ParseInt(parsedInput, 10, 64)
	if err != nil {
		return &Request{}, err
	}

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
