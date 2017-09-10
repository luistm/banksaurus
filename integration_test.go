package main

import (
	"go-cli-bank/categories"
	"go-cli-bank/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	categoryName := "TestCategory"
	dbHandler := infrastructure.DatabaseHandler{}
	cr := categories.CategoryRepository{DBHandler: &dbHandler}
	i := categories.Interactor{Repository: &cr}
	c, err := i.NewCategory(categoryName)
	assert.NoError(t, err)
	assert.Equal(t, categoryName, c.Name, "Fetches category")

	c, err = i.GetCategory(categoryName)
	assert.NoError(t, err)
	assert.Equal(t, categoryName, c.Name, "Fetches category")
}
