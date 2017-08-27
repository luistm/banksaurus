package categories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryRepositorySave(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	cr := CategoryRepository{}
	err := cr.Save(&Category{})
	assert.Error(t, err)
}
