package entities

import (
	"github.com/stretchr/testify/mock"
)

type MockSQLStorage struct {
	mock.Mock
}

func (m *MockSQLStorage) Execute(statement string, values ...interface{}) error {
	arguments := []interface{}{statement}
	arguments = append(arguments, values...)
	args := m.Called(arguments...)
	return args.Error(0)
}

func (m *MockSQLStorage) Query(statement string) (Row, error) {
	args := m.Called(statement)
	return args.Get(0).(Row), args.Error(1)
}
