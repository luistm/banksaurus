package lib

import (
	"github.com/stretchr/testify/mock"
)

// MockSQLStorage to use in tests which need an SQLDatabaseHandler
type MockSQLStorage struct {
	mock.Mock
}

// Execute method mock
func (m *MockSQLStorage) Execute(statement string, values ...interface{}) error {
	arguments := []interface{}{statement}
	arguments = append(arguments, values...)
	args := m.Called(arguments...)
	return args.Error(0)
}

// Query method mock
func (m *MockSQLStorage) Query(statement string, a ...interface{}) (Rows, error) {
	args := m.Called(statement, a)
	return args.Get(0).(Rows), args.Error(1)
}

// RepositoryMock to use in tests which require a Repository
type RepositoryMock struct {
	mock.Mock
}

// Save method mock
func (m *RepositoryMock) Save(c Entity) error {
	args := m.Called(c)
	return args.Error(0)
}

// Get method mock
func (m *RepositoryMock) Get(s string) (Entity, error) {
	args := m.Called(s)
	return args.Get(0).(Entity), args.Error(1)
}

// GetAll method mock
func (m *RepositoryMock) GetAll() ([]Entity, error) {
	args := m.Called()
	return args.Get(0).([]Entity), args.Error(1)
}
