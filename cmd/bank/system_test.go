package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystem(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

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
			name:          "Create category",
			command:       []string{"category", "new", "testCategory"},
			expected:      "Created category 'testCategory'",
			errorExpected: false,
		},
		{
			name:          "Shows report from file",
			command:       []string{"report", "--input", "../../tests/fixtures/test_file.csv"},
			expected:      "Expense is  0\nCredit is  0\n",
			errorExpected: false,
		},
	}

	for _, tc := range testCases {
		cmd := exec.Command("../../bank", tc.command...)

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
