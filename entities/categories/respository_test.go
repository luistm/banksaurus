package categories

import (
	"errors"
	"testing"

	"github.com/luistm/go-bank-cli/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockSQLStorage struct {
	mock.Mock
}

func (m *mockSQLStorage) Execute(statement string, value ...interface{}) error {
	args := m.Called(statement)
	return args.Error(0)
}

func (m *mockSQLStorage) Query(statement string) (entities.Row, error) {
	args := m.Called(statement)
	return args.Get(0).(entities.Row), args.Error(1)
}

func TestCategoryRepositoryGetAll(t *testing.T) {

	testCases := []struct {
		name string
	}{}

	for _, tc := range testCases {
		m := new(mockSQLStorage)
		cr := repository{SQLStorage: m}
		_, err := cr.GetAll()
		assert.Error(t, err, tc.name)
	}
}

func TestUnitCategoryRepositorySave(t *testing.T) {

	// Category has no name
	name := "Returns error if category has no name"
	m := new(mockSQLStorage)
	cr := repository{SQLStorage: m}
	err := cr.Save(&Category{})
	assert.EqualError(t, err, "Invalid category", name)

	// Category is nil
	name = "Returns error if category is nil"
	m = new(mockSQLStorage)
	cr = repository{SQLStorage: m}
	err = cr.Save(nil)
	assert.EqualError(t, err, "Invalid category", name)

	// Infrastructure failure
	name = "Returns error if it fails to save category into infrastructure"
	m = new(mockSQLStorage)
	m.On("Execute", insertStatement).Return(errors.New("TestError"))
	cr = repository{SQLStorage: m}
	c := Category{Name: "TestCategory"}
	err = cr.Save(&c)
	m.AssertExpectations(t)
	assert.Error(t, err)
	assert.EqualError(t, err, "Infrastructure error: TestError", name)

	// Success
	name = "Saves category to infrastructure"
	m = new(mockSQLStorage)
	m.On("Execute", insertStatement).Return(nil)
	cr = repository{SQLStorage: m}
	c = Category{Name: "TestCategory"}
	err = cr.Save(&c)
	m.AssertExpectations(t)
	assert.NoError(t, err)
}
