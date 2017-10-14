package categories

import (
	"errors"
	"testing"

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

func TestUnitGetCategories(t *testing.T) {

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
					errors.New("Repository mock error"),
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
		i := new(Interactor)
		if tc.mock != nil {
			i.Repository = tc.mock
			tc.mock.On(tc.mInput.method).Return(tc.mInput.returnArguments...)
		}

		cats, err := i.GetCategories()

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

func TestUnitGetCategory(t *testing.T) {

	m := new(repositoryMock)
	i := new(Interactor)

	i.Repository = m
	categoryName := "testCategory"

	name := "Fails to get the category due to repository failure"
	m.On("Get", categoryName).Return(&Category{}, errors.New("Error"))
	_, err := i.GetCategory(categoryName)
	m.AssertExpectations(t)
	assert.EqualError(t, err, "Failed to get category: Error", name)

	name = "Fails to get category when no repository available"
	i = new(Interactor)
	_, err = i.GetCategory(categoryName)
	assert.EqualError(t, err, "Repository is not defined", name)

	name = "Fails to get category if name is not defined"
	i = new(Interactor)
	i.Repository = m
	_, err = i.GetCategory("")
	assert.EqualError(t, err, "Cannot get category whitout a category name", name)

	name = "Gets specified category"
	i = new(Interactor)
	m = new(repositoryMock)
	i.Repository = m
	m.On("Get", categoryName).Return(&Category{Name: categoryName}, nil)
	c, err := i.GetCategory(categoryName)
	m.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(c), name)
	assert.Equal(t, categoryName, c[0].Name, name)
}

func TestUnitNewCategory(t *testing.T) {

	m := new(repositoryMock)
	i := new(Interactor)
	i.Repository = m
	categoryName := "testCategory"

	name := "Fails to create a new category due to repository failure"
	m.On("Save", &Category{Name: categoryName}).Return(errors.New("Error"))
	c, err := i.NewCategory(categoryName)
	m.AssertExpectations(t)
	assert.EqualError(t, err, "Failed to create category: Error", name)

	name = "Fails to create category if repository is not defined"
	i = new(Interactor)
	_, err = i.NewCategory(categoryName)
	assert.EqualError(t, err, "Repository is not defined", name)

	name = "Fails to create category is name is empty"
	i = new(Interactor)
	_, err = i.NewCategory("")
	assert.EqualError(t, err, "Cannot create category whitout a category name")

	name = "Creates category with specified name"
	i = new(Interactor)
	m = new(repositoryMock)
	i.Repository = m
	m.On("Save", &Category{Name: categoryName}).Return(nil)
	c, err = i.NewCategory(categoryName)
	m.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(c), name)
	assert.Equal(t, categoryName, c[0].Name, name)

}
