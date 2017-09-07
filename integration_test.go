package main

import (
	"expensetracker/categories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	categoryName := "TestCategory"
	CommandCreateCategory(categoryName)

	i := new(categories.Interactor)
	_, err := i.GetCategory(categoryName)
	assert.NoError(t, err)
	// assert.Equal(t, c.name, categoryName, "Fetches category")
}
