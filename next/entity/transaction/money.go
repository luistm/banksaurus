package transaction

import "errors"

// ErrInvalidMoneyValue ...
var ErrInvalidMoneyValue = errors.New("invalid money value")

// NewMoney creates an instance of money
func NewMoney(value uint64) (*Money, error){
	if value == 0 {
		return &Money{}, ErrInvalidMoneyValue
	}

	return  &Money{}, nil
}

// Money value object
type Money struct{}