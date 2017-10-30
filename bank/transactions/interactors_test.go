package transactions

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/lib/customerrors"
	"github.com/luistm/go-bank-cli/lib/sellers"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) GetAll() ([]*Transaction, error) {
	args := m.Called()
	return args.Get(0).([]*Transaction), args.Error(1)
}

func TestUnitInteractorTransactionsLoad(t *testing.T) {

	testCasesRepository := []struct {
		name       string
		output     error
		withMock   bool
		mockOutput []interface{}
	}{
		{
			name:       "Returns error if repository is not defined",
			output:     customerrors.ErrRepositoryUndefined,
			withMock:   false,
			mockOutput: nil,
		},
		{
			name:       "Returns error on repository error",
			output:     &customerrors.ErrRepository{Msg: "Test Error"},
			withMock:   true,
			mockOutput: []interface{}{[]*Transaction{}, errors.New("Test Error")},
		},
	}

	for _, tc := range testCasesRepository {
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

	t1 := &Transaction{s: sellers.New("d1", "")}
	t2 := &Transaction{s: sellers.New("d2", "")}
	i := interactor{}
	m := new(repositoryMock)
	i.repository = m
	m.On("GetAll").Return([]*Transaction{t1, t2})

	testCasesEntityCreator := []struct {
		name       string
		output     error
		mockOutput []interface{}
	}{
		{
			name:   "Returns error if entity creator returns fails",
			output: &customerrors.ErrInteractor{Msg: "Test Error"},
			// withMock: true
			mockOutput: []interface{}{t1.s, errors.New("Test Error")},
		},
	}

	for _, tc := range testCasesEntityCreator {

		err := i.Load()

		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
	}

}
