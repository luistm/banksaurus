package money

import (
	"errors"
	"strconv"
)

// ErrInvalidMoneyValue ...
var ErrInvalidMoneyValue = errors.New("invalid money value")

// NewMoney creates an instance of money
func NewMoney(value int64) (*Money, error) {
	if value == 0 {
		return &Money{}, ErrInvalidMoneyValue
	}

	return &Money{value}, nil
}

// Money value object
type Money struct{ value int64 }

// Add an instance of money an returns another representing the result
func (m *Money) Add(money *Money) (*Money, error) {
	// TODO: What happens if m.value was never defined?
	newValue := m.value + money.value
	return &Money{value: newValue}, nil
}

func (m *Money) Value() int64 {
	return m.value
}

// String to satisfy the fmt.Stringer
func (m *Money) String() string {
	s := strconv.FormatInt(m.value, 10)
	return s
}
