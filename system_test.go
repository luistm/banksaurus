package main

import (
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
			expected:      "Usage:\n\tgo-cli-bank -h | --help\n",
			errorExpected: true,
		},
		{
			name:          "Shows usage if option is '-h'",
			command:       []string{"-h"},
			expected:      usage + "\n",
			errorExpected: false,
		},
		{
			name:          "Create category",
			command:       []string{"new", "category", "testCategory"},
			expected:      "Created category 'testCategory'",
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		cmd := exec.Command("go-cli-bank", tc.command...)

		stdoutStderr, err := cmd.CombinedOutput()

		if tc.errorExpected {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, tc.expected, string(stdoutStderr), tc.name)
	}
}
