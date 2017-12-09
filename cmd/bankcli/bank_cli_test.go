package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/luistm/go-bank-cli/elib/testkit"
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
			expected:      "77.52 COMPRA CONTINENTE MAI\n95.09 COMPRA FARMACIA SAO J\n95.09 COMPRA FARMACIA SAO J",
			errorExpected: false,
		},
		{
			name:     "No sellers should be available here",
			command:  []string{"seller", "show"},
			expected: "",
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
			name:     "LoadDataFromRecords records from file",
			command:  []string{"load", "--input", "./tests/fixtures/sample_records_load.csv"},
			expected: "",
		},
		{
			name:     "Shows sellers loaded by the run report",
			command:  []string{"seller", "show"},
			expected: "COMPRA CONTINENTE MAI\nCOMPRA FARMACIA SAO J\n",
		},
		{
			name:     "Adds pretty name to seller",
			command:  []string{"seller", "change", "COMPRA CONTINENTE MAI", "--pretty", "Continente"},
			expected: "",
		},
		{
			name:     "Show seller changed",
			command:  []string{"seller", "show"},
			expected: "Continente\nCOMPRA FARMACIA SAO J\n",
		},
		{
			name:          "Shows report from bank records file, with sellers name instead of slug",
			command:       []string{"report", "--input", "./tests/fixtures/sample_records_load.csv"},
			expected:      "77.52 Continente\n95.09 COMPRA FARMACIA SAO J\n95.09 COMPRA FARMACIA SAO J",
			errorExpected: false,
		},
		{
			name: "Shows report from bank records file, grouped by seller",
			command: []string{
				"report",
				"--input", "./tests/fixtures/sample_records_load.csv",
				"--grouped",
			},
			expected:      "77.52 Continente\n190.18 COMPRA FARMACIA SAO J",
			errorExpected: false,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		t.Log(fmt.Sprintf("$ bankcli %s", strings.Join(tc.command, " ")))
		cmd := exec.Command("../../bankcli", tc.command...)
		stdoutStderr, err := cmd.CombinedOutput()

		testkit.AssertEqual(t, tc.expected, string(stdoutStderr))
		if !tc.errorExpected && err != nil {
			t.Fatalf("System test command failed: %s", err)
		}
	}

	// Remove any test files
	if err := os.RemoveAll(DatabasePath + "/" + DatabaseName + ".db"); err != nil {
		t.Error(err)
	}
}

func TestUnitNewCommand(t *testing.T) {

	testCases := []struct {
		name   string
		input  []string
		output []interface{}
	}{
		{
			name:   "Returns error if command not found",
			input:  []string{"thisCommandDoesNotExist"},
			output: []interface{}{&ReportCommand{}, errCommandNotFound},
		},
		{
			name:   "Returns error if empty slice received",
			input:  []string{},
			output: []interface{}{&ReportCommand{}, errCommandIsUndefined},
		},
		{
			name:   "Returns command instance if cli input matches",
			input:  []string{"report"},
			output: []interface{}{&ReportCommand{commandType: "report"}, nil},
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)

		command, err := newCommand(tc.input)

		got := []interface{}{command, err}
		testkit.AssertEqual(t, tc.output, got)
	}
}
