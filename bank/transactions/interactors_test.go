package transactions

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/entities"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) GetAll() ([]*record, error) {
	args := m.Called()
	return args.Get(0).([]*record), args.Error(1)
}

func TestUnitInteractorTransactionsLoad(t *testing.T) {

	testCases := []struct {
		name       string
		output     error
		withMock   bool
		mockOutput []interface{}
	}{
		{
			name:       "Returns error if repository is not defined",
			output:     entities.ErrRepositoryUndefined,
			withMock:   false,
			mockOutput: nil,
		},
		{
			name:       "Returns error on repository error",
			output:     &entities.ErrRepository{Msg: "Test Error"},
			withMock:   true,
			mockOutput: []interface{}{[]*record{}, errors.New("Test Error")},
		},
		// {
		// 	name: "Creates description",
		// }
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := interactor{}
		var m *repositoryMock
		if tc.withMock {
			m = new(repositoryMock)
			i.repository = m
			m.On("GetAll").Return(tc.mockOutput...)
		}

		err := i.Load()

		if tc.withMock {
			m.AssertExpectations(t)
		}
		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
	}
}
