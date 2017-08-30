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

func TestNewCategory(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	m := new(repositoryMock)
	i := new(Interactor)
	i.Repository = m
	categoryName := "testCategory"

	name := "Fails to create a new category due to repository failure"
	m.On("Save", &Category{name: categoryName}).Return(errors.New("Failed to create category due to repository failure"))
	c, err := i.NewCategory(categoryName)
	m.AssertExpectations(t)
	assert.Error(t, err)
	assert.Equal(t, &Category{}, c, name)

	name = "Fails to create category if repository is not defined"
	i = new(Interactor)
	_, err = i.NewCategory(categoryName)
	assert.Error(t, err)

	// name = "Fails to create category is name is empty"
	// c, err = NewCategory("")
	// m.AssertExpectations(t)
	// assert.Error(t, err)
	// assert.Equal(t, &Category{}, c, name)

	// name = "Creates category with specified name"
	// c, err = NewCategory(categoryName)
	// m.AssertExpectations(t)
	// assert.NoError(t, err)
	// assert.Equal(t, c.name, categoryName, name)

	// name = "Allows to recreate existing category"
	// c, err = NewCategory(categoryName)
	// m.AssertExpectations(t)
	// assert.NoError(t, err)
	// assert.Equal(t, c.name, categoryName, name)
}
