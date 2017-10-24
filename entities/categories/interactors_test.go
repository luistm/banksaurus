package categories

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luistm/go-bank-cli/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) Save(c *Category) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *repositoryMock) Get(s string) (*Category, error) {
	args := m.Called(s)
	return args.Get(0).(*Category), args.Error(1)
}

func (m *repositoryMock) GetAll() ([]*Category, error) {
	args := m.Called()
	return args.Get(0).([]*Category), args.Error(1)
}

func TestUnitGetAll(t *testing.T) {

	testCases := []struct {
		name          string
		expectedLen   int
		errorExpected bool
		mock          *repositoryMock
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
			mock:          new(repositoryMock),
			mInput: &struct {
				method          string
				returnArguments []interface{}
			}{
				method: "GetAll",
				returnArguments: []interface{}{
					[]*Category{},
					errors.New("repository mock error"),
				},
			},
		},
		{
			name:          "Returns slice of categories",
			expectedLen:   1,
			errorExpected: false,
			mock:          new(repositoryMock),
			mInput: &struct {
				method          string
				returnArguments []interface{}
			}{
				method: "GetAll",
				returnArguments: []interface{}{
					[]*Category{&Category{Name: "ThisIsATestCategory"}},
					nil,
				},
			},
		},
	}

	for _, tc := range testCases {
		i := new(interactor)
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

	categoryName := "Category Name"

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
				[]*Category{},
				entities.ErrRepositoryUndefined,
			},
			withMock:   false,
			mockInput:  "",
			mockOutput: nil,
		},
		{
			name:  "Returns empty result if categoryName name is not defined",
			input: "",
			output: []interface{}{
				[]*Category{},
				nil,
			},
			withMock:   false,
			mockInput:  "",
			mockOutput: nil,
		},
		{
			name:  "Returns error on respository error",
			input: categoryName,
			output: []interface{}{
				[]*Category{},
				&entities.ErrRepository{Msg: "Test Error"},
			},
			withMock:   true,
			mockInput:  "",
			mockOutput: []interface{}{&Category{}, errors.New("Test Error")},
		},
		{
			name:  "Returns list of categories with one categoryName",
			input: categoryName,
			output: []interface{}{
				[]*Category{&Category{Name: categoryName}},
				nil,
			},
			withMock:   true,
			mockInput:  categoryName,
			mockOutput: []interface{}{&Category{Name: categoryName}, nil},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &interactor{}
		var m *repositoryMock
		if tc.withMock {
			m = new(repositoryMock)
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
				[]*Category{},
				entities.ErrRepositoryUndefined,
			},
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:  "Returns empty list if input is empty",
			input: "",
			output: []interface{}{
				[]*Category{},
				entities.ErrRepositoryUndefined,
			},
			withMock:   false,
			mockInput:  nil,
			mockOutput: nil,
		},
		{
			name:  "Returns error on repository error",
			input: categoryNameName,
			output: []interface{}{
				[]*Category{},
				&entities.ErrRepository{Msg: "Test Error"},
			},
			withMock:   true,
			mockInput:  &Category{Name: categoryNameName},
			mockOutput: errors.New("Test Error"),
		},
		{
			name:  "Adds categoryName to repository",
			input: categoryNameName,
			output: []interface{}{
				[]*Category{&Category{Name: categoryNameName}},
				nil,
			},
			withMock:   true,
			mockInput:  &Category{Name: categoryNameName},
			mockOutput: nil,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &interactor{}
		var m *repositoryMock
		if tc.withMock {
			m = new(repositoryMock)
			m.On("Save", tc.mockInput).Return(tc.mockOutput)
			i.repository = m
		}

		c, err := i.Add(categoryNameName)

		if tc.withMock {
			m.AssertExpectations(t)
		}
		got := []interface{}{c, err}
		if !reflect.DeepEqual(tc.output, got) {
			t.Errorf("Expected '%v', got '%v'", tc.output, got)
		}
	}
}
