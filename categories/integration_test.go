package categories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	categoryName := "TestCategory"
	i := new(Interactor)

	c, err := i.NewCategory(categoryName)
	assert.NoError(t, err)
	assert.Equal(t, c.name, categoryName)

	// c, err = GetCategory(c.name)
	// assert.NoError(t, err)
	// assert.Equal(t, c.name, categoryName)
}
