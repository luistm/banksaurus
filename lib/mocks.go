package lib

import (
	"github.com/stretchr/testify/mock"
)

// SQLStorageMock to use in tests which need an SQLInfrastructer
type SQLStorageMock struct {
	mock.Mock
}

// Execute method mock
func (m *SQLStorageMock) Execute(statement string, values ...interface{}) error {
	arguments := []interface{}{statement}
	arguments = append(arguments, values...)
	args := m.Called(arguments...)
	return args.Error(0)
}

// Query method mock
func (m *SQLStorageMock) Query(statement string, a ...interface{}) (Rows, error) {
	args := m.Called(statement, a)
	return args.Get(0).(Rows), args.Error(1)
}

// RepositoryMock to use in tests which require a Repository
type RepositoryMock struct {
	mock.Mock
}

// Save method mock
func (m *RepositoryMock) Save(c Identifier) error {
	args := m.Called(c)
	return args.Error(0)
}

// Get method mock
func (m *RepositoryMock) Get(s string) (Identifier, error) {
	args := m.Called(s)
	return args.Get(0).(Identifier), args.Error(1)
}

// GetAll method mock
func (m *RepositoryMock) GetAll() ([]Identifier, error) {
	args := m.Called()
	return args.Get(0).([]Identifier), args.Error(1)
}

// PresenterMock to use in tests which need a presenter
type PresenterMock struct {
	mock.Mock
}

// Present ...
func (m *PresenterMock) Present(entities []Identifier) error {
	args := m.Called(entities)
	return args.Error(0)
}
