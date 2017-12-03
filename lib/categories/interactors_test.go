package categories

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/lib"
	"github.com/luistm/go-bank-cli/lib/customerrors"
	"github.com/stretchr/testify/assert"
)

func TestUnitGetAll(t *testing.T) {

	testCases := []struct {
		name          string
		expectedLen   int
		errorExpected bool
		mock          *lib.RepositoryMock
		mInput        *struct {
			method          string
			returnArguments []interface{}
		}
	}{
		{
			name:          "Fails to get categories if repository is not defined",
			expectedLen:   0,
			errorExpected: true,
			mock:          nil,
			mInput:        nil,
		},
		{
			name:          "Fails to get categories on repository error",
			expectedLen:   0,
			errorExpected: true,
			mock:          new(lib.RepositoryMock),
			mInput: &struct {
				method          string
				returnArguments []interface{}
			}{
				method: "GetAll",
				returnArguments: []interface{}{
					[]lib.Identifier{},
					errors.New("repository mock error"),
				},
			},
		},
		{
			name:          "Returns slice of categories",
			expectedLen:   1,
			errorExpected: false,
			mock:          new(lib.RepositoryMock),
			mInput: &struct {
				method          string
				returnArguments []interface{}
			}{
				method: "GetAll",
				returnArguments: []interface{}{
					[]lib.Identifier{&Category{name: "ThisIsATestCategory"}},
					nil,
				},
			},
		},
	}

	for _, tc := range testCases {
		i := new(Interactor)
		if tc.mock != nil {
			i.repository = tc.mock
			tc.mock.On(tc.mInput.method).Return(tc.mInput.returnArguments...)
		}

		cats, err := i.GetAll()

		if tc.mock != nil {
			tc.mock.AssertExpectations(t)
		}
		if tc.errorExpected {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, len(cats), tc.expectedLen, tc.name)
		}
	}

}

func TestUnitInteractorGetCategory(t *testing.T) {

	categoryName := "Category name"

	testCases := []struct {
		name       string
		input      string
		output     []interface{}
		withMock   bool
		mockInput  string
		mockOutput []interface{}
	}{
		{
			name:  "Returns error if repository is undefined",
			input: categoryName,
			output: []interface{}{
				[]lib.Identifier{},
				customerrors.ErrRepositoryUndefined,
			},
			withMock:   false,
			mockInput:  "",
			mockOutput: nil,
		},
		{
			name:  "Returns empty result if categoryName name is not defined",
			input: "",
			output: []interface{}{
				[]lib.Identifier{},
				nil,
			},
			withMock:   false,
			mockInput:  "",
			mockOutput: nil,
		},
		{
			name:  "Returns error on repository error",
			input: categoryName,
			output: []interface{}{
				[]lib.Identifier{},
				&customerrors.ErrRepository{Msg: "Test Error"},
			},
			withMock:   true,
			mockInput:  "",
			mockOutput: []interface{}{&Category{}, errors.New("Test Error")},
		},
		{
			name:  "Returns list of categories with one categoryName",
			input: categoryName,
			output: []interface{}{
				[]lib.Identifier{&Category{name: categoryName}},
				nil,
			},
			withMock:   true,
			mockInput:  categoryName,
			mockOutput: []interface{}{&Category{name: categoryName}, nil},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &Interactor{}
		var m *lib.RepositoryMock
		if tc.withMock {
			m = new(lib.RepositoryMock)
			m.On("Get", tc.input).Return(tc.mockOutput...)
			i.repository = m
		}

		c, err := i.GetCategory(tc.input)

		if tc.withMock {
			m.AssertExpectations(t)
		}
		got := []interface{}{c, err}
		if !reflect.DeepEqual(tc.output, got) {
			t.Errorf("Expected '%v', got '%v'", tc.output, got)
		}
	}
}

func TestUnitInteractorAdd(t *testing.T) {

	categoryNameName := "testCategory"

	testCases := []struct {
		name       string
		input      string
		output     []interface{}
		withMock   bool
		mockInput  *Category
		mockOutput error
	}{
		{
			name:  "Returns error if repository is not defined",
			input: categoryNameName,
			output: []interface{}{
				[]lib.Identifier{},
				customerrors.ErrRepositoryUndefined,
			},
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:  "Returns empty list if input is empty",
			input: "",
			output: []interface{}{
				[]lib.Identifier{},
				customerrors.ErrRepositoryUndefined,
			},
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:  "Returns error on repository error",
			input: categoryNameName,
			output: []interface{}{
				[]lib.Identifier{},
				&customerrors.ErrRepository{Msg: "Test Error"},
			},
			withMock:   true,
			mockInput:  &Category{name: categoryNameName},
			mockOutput: errors.New("Test Error"),
		},
		{
			name:  "Adds categoryName to repository",
			input: categoryNameName,
			output: []interface{}{
				[]lib.Identifier{&Category{name: categoryNameName}},
				nil,
			},
			withMock:   true,
			mockInput:  &Category{name: categoryNameName},
			mockOutput: nil,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &Interactor{}
		var m *lib.RepositoryMock
		if tc.withMock {
			m = new(lib.RepositoryMock)
			m.On("Save", tc.mockInput).Return(tc.mockOutput)
			i.repository = m
		}

		c, err := i.Create(categoryNameName)

		if tc.withMock {
			m.AssertExpectations(t)
		}
		got := []interface{}{c, err}
		if !reflect.DeepEqual(tc.output, got) {
			t.Errorf("Expected '%v', got '%v'", tc.output, got)
		}
	}
}
