package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	app "github.com/luistm/banksaurus/cmd/bscli/application"
	"github.com/luistm/testkit"
)

func deleteTestFiles(t *testing.T) {
	if err := os.RemoveAll(app.Path()); err != nil {
		t.Error(err)
	}
}

func TestMain(t *testing.M) {
	os.Setenv("BANKSAURUS_ENV", "dev")
	defer os.Setenv("BANKSAURUS_ENV", "")

	os.Exit(t.Run())
}

func TestAcceptanceUsage(t *testing.T) {

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
		t.Log(fmt.Sprintf("$ bscli %s", strings.Join(tc.command, " ")))
		cmd := exec.Command("../../bscli", tc.command...)
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

func TestAcceptance(t *testing.T) {

	defer deleteTestFiles(t)

	fixture := "./data/fixtures/sample_records_load.csv"

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
			expected:      app.Version + "\n",
			errorExpected: false,
		},
		{
			name:     "No seller should be available here",
			command:  []string{"seller", "show"},
			expected: "",
		},
		{
			name:     "Load records from file",
			command:  []string{"load", "--input", fixture},
			expected: "",
		},
		{
			name:     "Shows seller loaded by the load records from file",
			command:  []string{"seller", "show"},
			expected: "COMPRA CONTINENTE MAI\nTRF CREDIT\nCOMPRA FARMACIA SAO J\n",
		},
		{
			name:          "Shows report with all available transactions",
			command:       []string{"report"},
			expected:      "-0,52€  COMPRA CONTINENTE MAI\n593,48€ TRF CREDIT\n-95,09€ COMPRA FARMACIA SAO J\n-95,09€ COMPRA FARMACIA SAO J\n",
			errorExpected: false,
		},
		{
			name:          "Shows report from bank records file",
			command:       []string{"report", "--input", fixture},
			expected:      "-0,52€  COMPRA CONTINENTE MAI\n593,48€ TRF CREDIT\n-95,09€ COMPRA FARMACIA SAO J\n-95,09€ COMPRA FARMACIA SAO J\n",
			errorExpected: false,
		},
		{
			name:          "Shows report from bank records file, returns error if path does not exist",
			command:       []string{"report", "--input", "./thispathdoesnotexist/sample_records_load.csv"},
			expected:      errGeneric.Error() + "\n",
			errorExpected: true,
		},
		{
			name: "Shows report from bank records file, grouped by seller",
			command: []string{
				"report",
				"--input", fixture,
				"--grouped",
			},
			expected:      "-0,52€   COMPRA CONTINENTE MAI\n593,48€  TRF CREDIT\n-190,18€ COMPRA FARMACIA SAO J\n",
			errorExpected: false,
		},
		{
			name:     "Adds pretty name to seller",
			command:  []string{"seller", "change", "COMPRA CONTINENTE MAI", "--pretty", "Continente"},
			expected: "",
		},
		{
			name:     "Show seller changed",
			command:  []string{"seller", "show"},
			expected: "Continente\nTRF CREDIT\nCOMPRA FARMACIA SAO J\n",
		},
		{
			name:          "Shows report, with seller name instead of slug",
			command:       []string{"report"},
			expected:      "-0,52€  Continente\n593,48€ TRF CREDIT\n-95,09€ COMPRA FARMACIA SAO J\n-95,09€ COMPRA FARMACIA SAO J\n",
			errorExpected: false,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		t.Log(fmt.Sprintf("$ bscli %s", strings.Join(tc.command, " ")))
		cmd := exec.Command("../../bscli", tc.command...)
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
