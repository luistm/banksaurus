package app_test

import (
	"testing"

	"os"
	"path"

	"os/user"

	"github.com/luistm/banksaurus/app"
	"github.com/luistm/testkit"
)

func TestUnitNewApp(t *testing.T) {

	// Try to create app in existing directory, but invalid project path
	// Try to create app with non existing configuration file

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	testCase := []struct {
		name                  string
		configurationFilePath string
		expectedApp           *app.App
		expectedErr           string
	}{
		{
			name: "Creates app",
			configurationFilePath: path.Join(pwd, "..", "/configurations/test_conf.json"),
			expectedApp:           &app.App{ProjectPath: path.Join(pwd, "/..")},
		},
		{
			name: "Creates app with non exiting directory",
			configurationFilePath: "/ThisDirectoryDoesNotExist",
			expectedApp:           &app.App{},
			expectedErr:           app.ErrPathDoesNotExist.Error(),
		},
		{
			name: "Creates app with non exiting configuration file",
			configurationFilePath: path.Join(pwd, "..", "/configurations/t_conf.json"),
			expectedApp:           &app.App{},
			expectedErr:           app.ErrPathDoesNotExist.Error(),
		},
		{
			name: "Creates app in invalid project structure",
			configurationFilePath: "/tmp",
			expectedApp:           &app.App{},
			expectedErr:           app.ErrInvalidAppStructure.Error(),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			application, err := app.New(tc.configurationFilePath)

			if tc.expectedErr != "" {
				testkit.AssertEqual(t, tc.expectedErr, err.Error())
			}
			testkit.AssertEqual(t, tc.expectedApp, application)
		})
	}
}

func TestUnitValidatePath(t *testing.T) {

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	testCases := []struct {
		name   string
		input  string
		errMsg string
	}{
		{
			name:   "Returns error if path does not exist",
			input:  path.Join(pwd, "..", "/ThisPathDoesNotExist"),
			errMsg: app.ErrPathDoesNotExist.Error(),
		},
		//{
		//	name:   "Returns error if path is file",
		//	input:  path.Join(pwd, "..", "/app/app.go"),
		//	errMsg: app.ErrPathIsFile.Error(),
		//},
		{
			name:   "Returns error if path does not have prefix '/'",
			input:  "app/app.go",
			errMsg: app.ErrMalformedPath.Error(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.ValidatePath(tc.input)

			// Error is expected
			if tc.errMsg != "" {
				testkit.AssertEqual(t, tc.errMsg, err.Error())
			}
		})
	}
}

func TestUnitGetDataBasePath(t *testing.T) {

	usr, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	expectedDbName := "bank"
	expectedDbPath := path.Join(usr.HomeDir, ".bank")

	dbName, dbPath := app.DatabasePath()

	testkit.AssertEqual(t, expectedDbName, dbName)
	testkit.AssertEqual(t, expectedDbPath, dbPath)

	os.Setenv("BANKSAURUS_ENV", "dev")
	defer os.Setenv("BANKSAURUS_ENV", "")

	expectedDbName = "bank"
	expectedDbPath = os.TempDir()

	dbName, dbPath = app.DatabasePath()

	testkit.AssertEqual(t, expectedDbName, dbName)
	testkit.AssertEqual(t, expectedDbPath, dbPath)
}
