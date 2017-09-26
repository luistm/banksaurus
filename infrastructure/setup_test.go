package infrastructure

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitStorage(t *testing.T) {
	// NOTE: Actually i don't like tests which touch the disk
	// I will look into this when i have the time. For now it will do....

	if !testing.Short() {
		t.Skip()
	}

	testCases := []struct {
		name          string
		dbName        string
		dbPath        string
		errorExpected bool
	}{
		{
			name:          "Name is empty",
			dbName:        "",
			dbPath:        "ignoreThisForNow",
			errorExpected: true,
		},
		{
			name:          "Path is empty",
			dbName:        "ignoreThisForNow",
			dbPath:        "",
			errorExpected: true,
		},
		{
			name:          "Path does not exist",
			dbName:        "ignoreThisForNow",
			dbPath:        "./ThisPathDoesNotExist",
			errorExpected: false,
		},
	}

	for _, tc := range testCases {
		err := InitStorage(tc.dbName, tc.dbPath)

		if tc.errorExpected {
			assert.Error(t, err, tc.name)
		} else {
			assert.NoError(t, err, tc.name)

			// Remove any test files
			if err := os.RemoveAll(tc.dbPath); err != nil {
				t.Error(err)
			}
		}
	}

}
