package main

import (
	"go-cli-bank/categories"
	"go-cli-bank/infrastructure"
	"os/exec"
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

func TestSystem(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testCases := []struct {
		name     string
		command  []string
		expected string
	}{
		{
			name:     "No output if no option is defined",
			command:  []string{""},
			expected: "",
		},
		{
			name:     "Create category",
			command:  []string{"new", "category", "testCategory"},
			expected: "Created category 'testCategory'",
		},
	}

	for _, tc := range testCases {
		cmd := exec.Command("go-cli-bank", tc.command...)

		stdoutStderr, err := cmd.CombinedOutput()

		assert.NoError(t, err)
		assert.Equal(t, tc.expected, string(stdoutStderr), tc.name)
	}
}
