package descriptions

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/entities"
)

func TestUnitInteractorAdd(t *testing.T) {

	var description = "TestDescrition"

	testCases := []struct {
		name       string
		input      string
		output     []interface{}
		withMock   bool
		mockInput  *Description
		mockOutput error
	}{
		{
			name:       "Returns error if repository is not defined",
			input:      description,
			output:     []interface{}{&Description{}, entities.ErrRepositoryUndefined},
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:       "Returns error if description is empty string",
			input:      "",
			output:     []interface{}{&Description{}, entities.ErrBadInput},
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:  "Returns error on repository error",
			input: description,
			output: []interface{}{
				&Description{},
				&entities.ErrRepository{Msg: "Test Error"},
			},
			withMock:   true,
			mockInput:  &Description{slug: description},
			mockOutput: errors.New("Test Error"),
		},
		{
			name:  "Returns description entity created",
			input: description,
			output: []interface{}{
				&Description{slug: description},
				nil,
			},
			withMock:   true,
			mockInput:  &Description{slug: description},
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

		d, err := i.Create(tc.input)

		if tc.withMock {
			m.AssertExpectations(t)
		}
		got := []interface{}{d, err}
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
			name: "Returns description entities",
			output: []interface{}{
				[]entities.Entity{&Description{}, &Description{}},
				nil,
			},
			withMock: true,
			mockOutput: []interface{}{
				[]entities.Entity{&Description{}, &Description{}},
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

		descriptions, err := i.GetAll()

		if tc.withMock {
			m.AssertExpectations(t)
		}
		got := []interface{}{descriptions, err}
		if !reflect.DeepEqual(tc.output, got) {
			t.Errorf("Expected '%v', got '%v'", tc.output, got)
		}
	}
}
