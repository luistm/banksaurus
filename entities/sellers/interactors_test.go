package sellers

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/entities"
)

func TestUnitInteractorAdd(t *testing.T) {

	var seller = "TestDescrition"

	testCases := []struct {
		name       string
		input      string
		output     []interface{}
		withMock   bool
		mockInput  *Seller
		mockOutput error
	}{
		{
			name:       "Returns error if repository is not defined",
			input:      seller,
			output:     []interface{}{&Seller{}, entities.ErrRepositoryUndefined},
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:       "Returns error if seller is empty string",
			input:      "",
			output:     []interface{}{&Seller{}, entities.ErrBadInput},
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:  "Returns error on repository error",
			input: seller,
			output: []interface{}{
				&Seller{},
				&entities.ErrRepository{Msg: "Test Error"},
			},
			withMock:   true,
			mockInput:  &Seller{slug: seller},
			mockOutput: errors.New("Test Error"),
		},
		{
			name:  "Returns seller entity created",
			input: seller,
			output: []interface{}{
				&Seller{slug: seller},
				nil,
			},
			withMock:   true,
			mockInput:  &Seller{slug: seller},
			mockOutput: nil,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &interactor{}
		var m *entities.RepositoryMock
		if tc.withMock {
			m = new(entities.RepositoryMock)
			m.On("Save", tc.mockInput).Return(tc.mockOutput)
			i.repository = m
		}

		s, err := i.Create(tc.input)

		if tc.withMock {
			m.AssertExpectations(t)
		}
		got := []interface{}{s, err}
		if !reflect.DeepEqual(tc.output, got) {
			t.Errorf("Expected %v, got %v", tc.output, got)
		}

	}
}

func TestUnitInteractorGetAll(t *testing.T) {

	// testCategory := &Category{name: "Test Category "}

	testCases := []struct {
		name       string
		output     []interface{}
		withMock   bool
		mockOutput []interface{}
	}{
		{
			name:       "Returns error if repository is undefined",
			output:     []interface{}{[]entities.Entity{}, entities.ErrRepositoryUndefined},
			withMock:   false,
			mockOutput: nil,
		},
		{
			name:     "Returns error on respository error",
			output:   []interface{}{[]entities.Entity{}, &entities.ErrRepository{Msg: "Test Error"}},
			withMock: true,
			mockOutput: []interface{}{
				[]entities.Entity{},
				errors.New("Test Error"),
			},
		},
		{
			name: "Returns seller entities",
			output: []interface{}{
				[]entities.Entity{&Seller{}, &Seller{}},
				nil,
			},
			withMock: true,
			mockOutput: []interface{}{
				[]entities.Entity{&Seller{}, &Seller{}},
				nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := interactor{}
		var m *entities.RepositoryMock
		if tc.withMock {
			m = new(entities.RepositoryMock)
			m.On("GetAll").Return(tc.mockOutput...)
			i.repository = m
		}

		sellers, err := i.GetAll()

		if tc.withMock {
			m.AssertExpectations(t)
		}
		got := []interface{}{sellers, err}
		if !reflect.DeepEqual(tc.output, got) {
			t.Errorf("Expected '%v', got '%v'", tc.output, got)
		}
	}
}
