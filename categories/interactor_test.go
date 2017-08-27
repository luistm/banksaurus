package categories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	name := "testCategory"
	err := NewCategory(name)
	assert.NoError(t, err)
}
