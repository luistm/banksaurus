package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystem(t *testing.T) {

	testCases := []struct {
		name          string
		command       []string
		expected      string
		errorExpected bool
	}{
		{
			name:          "Shows usage if no option is defined",
			command:       []string{""},
			expected:      usage + "\n",
			errorExpected: true,
		},
		{
			name:          "Shows usage if option is '-h'",
			command:       []string{"-h"},
			expected:      intro + usage + options + "\n",
			errorExpected: false,
		},
		{
			name:          "Shows report from file",
			command:       []string{"report", "--input", "./tests/fixtures/test_file.csv"},
			expected:      "Expense is  0\nCredit is  0\n",
			errorExpected: false,
		},
	}

	for _, tc := range testCases {
		cmd := exec.Command("../../bankcli", tc.command...)

		stdoutStderr, err := cmd.CombinedOutput()

		if tc.errorExpected {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)

			// Remove any test files
			if err := os.RemoveAll(DatabasePath + "/" + DatabaseName + ".db"); err != nil {
				t.Error(err)
			}

		}
		assert.Equal(t, tc.expected, string(stdoutStderr), tc.name)
	}
}

func TestSystemSellers(t *testing.T) {

	name := "Shows report from file"
	command := []string{"report", "--input", "./tests/fixtures/test_file.csv"}
	expected := "Expense is  0\nCredit is  0\n"

	cmd := exec.Command("../../bankcli", command...)
	stdoutStderr, err := cmd.CombinedOutput()

	assert.NoError(t, err)
	assert.Equal(t, expected, string(stdoutStderr), name)

	name = "Shows sellers loaded by the run report"
	command = []string{"seller", "show"}
	expected = "COMPRA CONTINENTE MAI\nLEVANTAMENTO Est Circ\nCOMPRA MODELO BONJOUR\nCOMPRA LIDL   CIA  LJ\nBX VALOR 03 TRANSACCO\n"

	cmd = exec.Command("../../bankcli", command...)
	stdoutStderr, err = cmd.CombinedOutput()

	assert.NoError(t, err)
	assert.Equal(t, expected, string(stdoutStderr), name)

	// Remove any test files
	if err := os.RemoveAll(DatabasePath + "/" + DatabaseName + ".db"); err != nil {
		t.Error(err)
	}

}

func TestSystemCategories(t *testing.T) {

	categoryName := "ThisIsACategoryNameForTesting"

	name := "Create category"
	command := []string{"category", "new", categoryName}
	expected := "Created category '" + categoryName + "'"

	cmd := exec.Command("../../bankcli", command...)
	stdoutStderr, err := cmd.CombinedOutput()

	assert.NoError(t, err)
	assert.Equal(t, expected, string(stdoutStderr), name)

	name = "Show category"
	command = []string{"category", "show"}
	expected = categoryName + "\n"

	cmd = exec.Command("../../bankcli", command...)
	stdoutStderr, err = cmd.CombinedOutput()

	assert.NoError(t, err)
	assert.Equal(t, expected, string(stdoutStderr), name)

	if err := os.RemoveAll(DatabasePath + "/" + DatabaseName + ".db"); err != nil {
		t.Error(err)
	}
}
