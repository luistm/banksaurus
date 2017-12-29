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
			name:     "Returns no error on success",
			expected: nil,
			mock:     new(lib.RepositoryMock),
			mockOutput: []interface{}{
				[]lib.Identifier{&Category{name: "ThisIsATestCategory"}},
				nil,
			},
		},
	}

	presenterMock := new(lib.PresenterMock)
	presenterMock.On("Present", []lib.Identifier{&Category{name: "ThisIsATestCategory"}}).Return(nil)

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &Interactor{presenter: presenterMock}
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

	testCasesPresenter := []struct {
		name       string
		output     error
		withMock   bool
		mockInput  []lib.Identifier
		mockOutput error
	}{
		{
			name:   "Returns error if presenter is not defined",
			output: customerrors.ErrPresenterUndefined,
		},
		{
			name:       "Returns error if presenter returns error",
			output:     &customerrors.ErrPresenter{Msg: "test error"},
			withMock:   true,
			mockInput:  []lib.Identifier{&Category{name: "ThisIsATestCategory"}},
			mockOutput: errors.New("test error"),
		},
	}

	repositoryMock := new(lib.RepositoryMock)
	repositoryMock.On("GetAll").Return([]lib.Identifier{&Category{name: "ThisIsATestCategory"}}, nil)

	for _, tc := range testCasesPresenter {
		t.Log(tc.name)
		i := &Interactor{repository: repositoryMock}
		var presenterMock *lib.PresenterMock
		if tc.withMock {
			presenterMock = new(lib.PresenterMock)
			presenterMock.On("Present", tc.mockInput).Return(tc.mockOutput)
			i.presenter = presenterMock
		}

		err := i.GetAll()

		if tc.withMock {
			presenterMock.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, err)
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

	presenterMock := new(lib.PresenterMock)
	presenterMock.On("Present", []lib.Identifier{}).Return(nil)

	for _, tc := range testCases {
		t.Log(tc.name)
		i := &Interactor{presenter: presenterMock}
		var repositoryMock *lib.RepositoryMock
		if tc.withMock {
			repositoryMock = new(lib.RepositoryMock)
			repositoryMock.On("Get", tc.input).Return(tc.mockOutput...)
			i.repository = repositoryMock
		}

		err := i.GetCategory(tc.input)

		if tc.withMock {
			repositoryMock.AssertExpectations(t)
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
		var repositoryMock *lib.RepositoryMock
		if tc.withMock {
			repositoryMock = new(lib.RepositoryMock)
			repositoryMock.On("Save", tc.mockInput).Return(tc.mockOutput)
			i.repository = repositoryMock
		}

		err := i.Create(tc.input)

		if tc.withMock {
			repositoryMock.AssertExpectations(t)
		}
		testkit.AssertEqual(t, tc.output, err)
	}
}
