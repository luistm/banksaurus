package bankservices

import (
	"github.com/luistm/banksaurus/banklib"
	"github.com/stretchr/testify/mock"
)

// Interactor is the interface each use case must implement
type Interactor interface {
	Execute() error
}

// Presenter is use
type Presenter interface {
	Present(...banklib.Entity) error
}

// PresenterMock to use in tests which need a presenter
type PresenterMock struct {
	mock.Mock
}

// Present ...
func (m *PresenterMock) Present(entities ...banklib.Entity) error {
	args := m.Called(entities)
	return args.Error(0)
}
