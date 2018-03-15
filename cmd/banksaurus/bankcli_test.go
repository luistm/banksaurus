package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/elib/testkit"
)

func deleteTestFiles(t *testing.T) {
	dbName, dbPath := configurations.DatabasePath()
	if err := os.RemoveAll(path.Join(dbPath, dbName) + ".db"); err != nil {
		t.Error(err)
	}
}

func TestSystemUsage(t *testing.T) {

	os.Setenv("GO_BANK_CLI_DEV", "true")
	defer os.Setenv("GO_BANK_CLI_DEV", "")
	defer deleteTestFiles(t)

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
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		t.Log(fmt.Sprintf("$ banksaurus %s", strings.Join(tc.command, " ")))
		cmd := exec.Command("../../banksaurus", tc.command...)
		var outBuffer, errBuffer bytes.Buffer
		cmd.Stdout = &outBuffer
		cmd.Stderr = &errBuffer

		err := cmd.Run()

		if !tc.errorExpected && err != nil {
			t.Log(outBuffer.String())
			t.Log(errBuffer.String())
			t.Fatalf("Test failed due to command error: %s", err.Error())
		}
		testkit.AssertEqual(t, tc.expected, errBuffer.String())
		testkit.AssertEqual(t, "", outBuffer.String())
	}
}

func TestSystem(t *testing.T) {

	os.Setenv("GO_BANK_CLI_DEV", "true")
	defer os.Setenv("GO_BANK_CLI_DEV", "")
	defer deleteTestFiles(t)

	testCases := []struct {
		name          string
		command       []string
		expected      string
		errorExpected bool
	}{
		{
			name:          "Shows usage if option is '-h'",
			command:       []string{"-h"},
			expected:      intro + usage + options + "\n",
			errorExpected: false,
		},
		{
			name:          "Shows version if option is '--version'",
			command:       []string{"--version"},
			expected:      version + "\n",
			errorExpected: false,
		},
		{
			name:          "Shows report from bank records file",
			command:       []string{"report", "--input", "./tests/fixtures/sample_records_load.csv"},
			expected:      "77.52 COMPRA CONTINENTE MAI\n95.09 COMPRA FARMACIA SAO J\n95.09 COMPRA FARMACIA SAO J\n",
			errorExpected: false,
		},
		{
			name:     "No sellers should be available here",
			command:  []string{"seller", "show"},
			expected: "",
		},
		// {
		// 	name:          "Create category",
		// 	command:       []string{"category", "new", "ThisIsACategoryNameForTesting"},
		// 	expected:      "",
		// 	errorExpected: false,
		// },
		// {
		// 	name:     "Show Category",
		// 	command:  []string{"category", "show"},
		// 	expected: "ThisIsACategoryNameForTesting\n",
		// },
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
			expected:      "77.52 Continente\n95.09 COMPRA FARMACIA SAO J\n95.09 COMPRA FARMACIA SAO J\n",
			errorExpected: false,
		},
		{
			name:          "Shows report from bank records file, with sellers name instead of slug",
			command:       []string{"report", "--input", "./thispathdoesnotexist/sample_records_load.csv"},
			expected:      errGeneric.Error() + "\n",
			errorExpected: true,
		},
		{
			name: "Shows report from bank records file, grouped by seller",
			command: []string{
				"report",
				"--input", "./tests/fixtures/sample_records_load.csv",
				"--grouped",
			},
			expected:      "77.52  Continente\n190.18 COMPRA FARMACIA SAO J\n",
			errorExpected: false,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		t.Log(fmt.Sprintf("$ banksaurus %s", strings.Join(tc.command, " ")))
		cmd := exec.Command("../../banksaurus", tc.command...)
		var outBuffer, errBuffer bytes.Buffer
		cmd.Stdout = &outBuffer
		cmd.Stderr = &errBuffer

		err := cmd.Run()

		if !tc.errorExpected && err != nil {
			t.Log(outBuffer.String())
			t.Log(errBuffer.String())
			t.Fatalf("Test failed due to command error: %s", err.Error())
		} else {
			if tc.errorExpected {
				testkit.AssertEqual(t, tc.expected, errBuffer.String())
				testkit.AssertEqual(t, "", outBuffer.String())
			} else {
				testkit.AssertEqual(t, "", errBuffer.String())
				testkit.AssertEqual(t, tc.expected, outBuffer.String())
			}
		}
	}
}
