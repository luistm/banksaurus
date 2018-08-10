package services

import (
	"github.com/luistm/banksaurus/lib"
	"github.com/stretchr/testify/mock"
)

// Servicer is the interface each use case must implement
type Servicer interface {
	Execute() error
}

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
