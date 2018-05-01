package bank

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/stretchr/testify/mock"
)

// Presenter is use
type Presenter interface {
	Present(...lib.Entity) error
}

// PresenterMock to use in tests which need a presenter
type PresenterMock struct {
	mock.Mock
}

// Present ...
func (m *PresenterMock) Present(entities ...lib.Entity) error {
	args := m.Called(entities)
	return args.Error(0)
}