package categories

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDBHandler struct {
	mock.Mock
}

func (m *mockDBHandler) Execute(statement string, value ...interface{}) error {
	args := m.Called(statement)
	return args.Error(0)
}

func (m *mockDBHandler) Query(statement string) (IRow, error) {
	args := m.Called(statement)
	return args.Get(0).(IRow), args.Error(1)
}

func TestCategoryRepositorySave(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	// Category has no name
	name := "Returns error if category has no name"
	m := new(mockDBHandler)
	cr := CategoryRepository{DBHandler: m}
	err := cr.Save(&Category{})
	assert.EqualError(t, err, "Invalid category", name)

	// Category is nil
	name = "Returns error if category is nil"
	m = new(mockDBHandler)
	cr = CategoryRepository{DBHandler: m}
	err = cr.Save(nil)
	assert.EqualError(t, err, "Invalid category", name)

	// Infrastructure failure
	name = "Returns error if it fails to save category into infrastructure"
	m = new(mockDBHandler)
	m.On("Execute", insertStatement).Return(errors.New("TestError"))
	cr = CategoryRepository{DBHandler: m}
	c := Category{Name: "TestCategory"}
	err = cr.Save(&c)
	m.AssertExpectations(t)
	assert.Error(t, err)
	assert.EqualError(t, err, "Infrastructure error: TestError", name)

	// Success
	name = "Saves category to infrastructure"
	m = new(mockDBHandler)
	m.On("Execute", insertStatement).Return(nil)
	cr = CategoryRepository{DBHandler: m}
	c = Category{Name: "TestCategory"}
	err = cr.Save(&c)
	m.AssertExpectations(t)
	assert.NoError(t, err)
}
