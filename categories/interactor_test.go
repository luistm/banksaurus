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
	_, err := NewCategory(name)
	assert.Error(t, err)
}

func TestGetCategory(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	name := "testCategory"
	_, err := GetCategory(name)
	assert.Error(t, err)
}
