package categories

import (
	"errors"
	"testing"

	"github.com/luistm/go-bank-cli/entities"
	"github.com/stretchr/testify/assert"
)

func TestCategoryRepositoryGetAll(t *testing.T) {

	testCases := []struct {
		name string
	}{}

	for _, tc := range testCases {
		m := new(entities.MockSQLStorage)
		cr := repository{SQLStorage: m}
		_, err := cr.GetAll()
		assert.Error(t, err, tc.name)
	}
}

func TestUnitCategoryRepositorySave(t *testing.T) {

	// Category has no name
	name := "Returns error if category has no name"
	m := new(entities.MockSQLStorage)
	cr := repository{SQLStorage: m}
	err := cr.Save(&Category{})
	assert.EqualError(t, err, "Invalid category", name)

	// Category is nil
	name = "Returns error if category is nil"
	m = new(entities.MockSQLStorage)
	cr = repository{SQLStorage: m}
	err = cr.Save(nil)
	assert.EqualError(t, err, "Invalid category", name)

	// Infrastructure failure
	name = "Returns error if it fails to save category into infrastructure"
	m = new(entities.MockSQLStorage)
	m.On("Execute", insertStatement).Return(errors.New("TestError"))
	cr = repository{SQLStorage: m}
	c := Category{Name: "TestCategory"}
	err = cr.Save(&c)
	m.AssertExpectations(t)
	assert.Error(t, err)
	assert.EqualError(t, err, "Infrastructure error: TestError", name)

	// Success
	name = "Saves category to infrastructure"
	m = new(entities.MockSQLStorage)
	m.On("Execute", insertStatement).Return(nil)
	cr = repository{SQLStorage: m}
	c = Category{Name: "TestCategory"}
	err = cr.Save(&c)
	m.AssertExpectations(t)
	assert.NoError(t, err)
}
