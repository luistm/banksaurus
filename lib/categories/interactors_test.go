package categories

import (
	"errors"
	"testing"

	"github.com/luistm/go-bank-cli/elib/testkit"

	"github.com/luistm/go-bank-cli/lib"
	"github.com/luistm/go-bank-cli/lib/customerrors"
)

func TestUnitGetAll(t *testing.T) {

	testCases := []struct {
		name       string
		expected   error
		mock       *lib.RepositoryMock
		mockOutput []interface{}
	}{
		{
			name:     "Returns error if repository is not defined",
			expected: customerrors.ErrRepositoryUndefined,
		},
		{
			name:       "Returns error if repository returns error",
			expected:   &customerrors.ErrRepository{Msg: "repository mock error"},
			mock:       new(lib.RepositoryMock),
			mockOutput: []interface{}{[]lib.Identifier{}, errors.New("repository mock error")},
		},
		{
			name:     "Returns not error on success",
			expected: nil,
			mock:     new(lib.RepositoryMock),
			mockOutput: []interface{}{
				[]lib.Identifier{&Category{name: "ThisIsATestCategory"}},
				nil,
			},
		},
	}

	for _, tc := range testCases {
		i := new(Interactor)
		if tc.mock != nil {
			i.repository = tc.mock
			tc.mock.On("GetAll").Return(tc.mockOutput...)
		}

		err := i.GetAll()

		if tc.mock != nil {
			tc.mock.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.expected, err)
	}

}

func TestUnitInteractorGetCategory(t *testing.T) {

	categoryName := "Category name"

	testCases := []struct {
		name       string
		input      string
		output     error
		withMock   bool
		mockInput  string
		mockOutput []interface{}
	}{
		{
			name:       "Returns error if repository is undefined",
			input:      categoryName,
			output:     customerrors.ErrRepositoryUndefined,
			withMock:   false,
			mockInput:  "",
			mockOutput: nil,
		},
		{
			name:     "Returns error if category name is not defined",
			output:   customerrors.ErrBadInput,
			withMock: false,
		},
		{
			name:       "Returns error on repository error",
			input:      categoryName,
			output:     &customerrors.ErrRepository{Msg: "Test Error"},
			withMock:   true,
			mockInput:  "",
			mockOutput: []interface{}{&Category{}, errors.New("Test Error")},
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

		err := i.GetCategory(tc.input)

		if tc.withMock {
			m.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, err)
	}
}

func TestUnitInteractorAdd(t *testing.T) {

	categoryNameName := "testCategory"

	testCases := []struct {
		name       string
		input      string
		output     error
		withMock   bool
		mockInput  *Category
		mockOutput error
	}{
		{
			name:   "Returns error if input is empty",
			output: customerrors.ErrBadInput,
		},
		{
			name:   "Returns error if repository is not defined",
			input:  categoryNameName,
			output: customerrors.ErrRepositoryUndefined,
		},
		{
			name:       "Returns error on repository error",
			input:      categoryNameName,
			output:     &customerrors.ErrRepository{Msg: "Test Error"},
			withMock:   true,
			mockInput:  &Category{name: categoryNameName},
			mockOutput: errors.New("Test Error"),
		},
		{
			name:      "Adds categoryName to repository",
			input:     categoryNameName,
			withMock:  true,
			mockInput: &Category{name: categoryNameName},
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

		err := i.Create(tc.input)

		if tc.withMock {
			m.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, err)
	}
}
