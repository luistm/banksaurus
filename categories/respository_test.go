package categories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDBHandler struct {
	mock.Mock
}

func (m *mockDBHandler) Execute(statement string) error {
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

	name := "Returns error if category has no name"
	m := new(mockDBHandler)
	cr := CategoryRepository{dbHandler: m}
	err := cr.Save(&Category{})
	assert.EqualError(t, err, "Invalid category", name)

	// Category is nil
	name = "Returns error if category is nil"
	m = new(mockDBHandler)
	cr = CategoryRepository{dbHandler: m}
	err = cr.Save(nil)
	assert.NoError(t, err)
}
