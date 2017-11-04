package transactions

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/lib/categories"

	"github.com/luistm/go-bank-cli/lib"

	"github.com/luistm/go-bank-cli/lib/customerrors"
	"github.com/luistm/go-bank-cli/lib/sellers"
	"github.com/stretchr/testify/mock"
)

type testMock struct {
	mock.Mock
}

func (m *testMock) GetAll() ([]*Transaction, error) {
	args := m.Called()
	return args.Get(0).([]*Transaction), args.Error(1)
}

func (m *testMock) Create(s string) (lib.Entity, error) {
	args := m.Called(s)
	return args.Get(0).(lib.Entity), args.Error(1)
}

func TestUnitInteractorTransactionsLoadDataFromRecords(t *testing.T) {

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
		i := Interactor{}
		var m *testMock
		if tc.withMock {
			m = new(testMock)
			i.repository = m
			m.On("GetAll").Return(tc.mockOutput...)
		}

		err := i.LoadDataFromRecords()

		if tc.withMock {
			m.AssertExpectations(t)
		}
		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
	}

	c := categories.New("Test Category")
	t1 := &Transaction{s: sellers.New("d1", "d1"), c: c}
	t2 := &Transaction{s: sellers.New("d2", "d2"), c: c}
	i := Interactor{}
	m := new(testMock)
	i.repository = m
	m.On("GetAll").Return([]*Transaction{t1, t2}, nil)

	testCasesEntityCreator := []struct {
		name       string
		output     error
		withMock   bool
		mockInput  string
		mockOutput []interface{}
	}{
		{
			name:       "Returns error if entity creator is not defined",
			output:     customerrors.ErrInteractorUndefined,
			withMock:   false,
			mockInput:  "",
			mockOutput: nil,
		},
		{
			name:       "Returns error if entity creator returns fails",
			output:     &customerrors.ErrInteractor{Msg: "Test Error"},
			withMock:   true,
			mockInput:  t1.s.String(),
			mockOutput: []interface{}{&sellers.Seller{}, errors.New("Test Error")},
		},
	}

	for _, tc := range testCasesEntityCreator {
		t.Log(tc.name)
		var im *testMock
		if tc.withMock {
			im = new(testMock)
			im.On("Create", tc.mockInput).Return(tc.mockOutput...)
			i.sellerInteractor = im
		}

		err := i.LoadDataFromRecords()

		if tc.withMock {
			im.AssertExpectations(t)
		}
		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
	}

	var sim *testMock
	sim = new(testMock)
	sim.On("Create", t1.s.String()).Return(t1.s, nil).
		On("Create", t2.s.String()).Return(t2.s, nil)
	i.sellerInteractor = sim

}
