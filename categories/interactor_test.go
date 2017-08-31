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

func TestGetCategory(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

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

}

func TestNewCategory(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	m := new(repositoryMock)
	i := new(Interactor)
	i.Repository = m
	categoryName := "testCategory"

	name := "Fails to create a new category due to repository failure"
	m.On("Save", &Category{name: categoryName}).Return(errors.New("Error"))
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
	m.On("Save", &Category{name: categoryName}).Return(nil)
	c, err = i.NewCategory(categoryName)
	m.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, categoryName, c.name, name)

}
