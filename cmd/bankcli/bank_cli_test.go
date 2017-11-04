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
			name:          "Shows report from bank records file",
			command:       []string{"report", "--input", "./tests/fixtures/sample_records_load.csv"},
			expected:      "77.52 COMPRA CONTINENTE MAI \n95.09 COMPRA FARMACIA SAO J ",
			errorExpected: false,
		},
		{
			name:          "Create category",
			command:       []string{"category", "new", "ThisIsACategoryNameForTesting"},
			expected:      "Created category 'ThisIsACategoryNameForTesting'",
			errorExpected: false,
		},
		{
			name:     "Show Category",
			command:  []string{"category", "show"},
			expected: "ThisIsACategoryNameForTesting\n",
		},
		{
			name:     "Loads records from file",
			command:  []string{"load", "--input", "./tests/fixtures/sample_records_load.csv"},
			expected: "",
		},
		{
			name:     "Shows sellers loaded by the run report",
			command:  []string{"seller", "show"},
			expected: "COMPRA CONTINENTE MAI \nCOMPRA FARMACIA SAO J \n",
		},
		{
			name:     "Adds pretty name to seller",
			command:  []string{"seller", "change", "COMPRA CONTINENTE MAI ", "--pretty", "Continente"},
			expected: "",
		},
		{
			name:     "Show seller changed",
			command:  []string{"seller", "show"},
			expected: "Continente\nCOMPRA FARMACIA SAO J \n",
		},
	}

	for _, tc := range testCases {
		cmd := exec.Command("../../bankcli", tc.command...)
		stdoutStderr, err := cmd.CombinedOutput()

		if tc.errorExpected {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, tc.expected, string(stdoutStderr), tc.name)
	}

	// Remove any test files
	if err := os.RemoveAll(DatabasePath + "/" + DatabaseName + ".db"); err != nil {
		t.Error(err)
	}
}
