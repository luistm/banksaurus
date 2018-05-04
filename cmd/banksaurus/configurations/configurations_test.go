package configurations

import (
	"os"
	"os/user"
	"path"
	"testing"

	"github.com/luistm/banksaurus/elib/testkit"
)

func TestUnitGetDataBasePath(t *testing.T) {

	usr, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	expectedDbName := "bankservices"
	expectedDbPath := path.Join(usr.HomeDir, ".bankservices")

	dbName, dbPath := DatabasePath()

	testkit.AssertEqual(t, expectedDbName, dbName)
	testkit.AssertEqual(t, expectedDbPath, dbPath)

	os.Setenv("GO_BANK_CLI_DEV", "true")
	defer os.Setenv("GO_BANK_CLI_DEV", "")

	expectedDbName = "bankservices"
	expectedDbPath = os.TempDir()

	dbName, dbPath = DatabasePath()

	testkit.AssertEqual(t, expectedDbName, dbName)
	testkit.AssertEqual(t, expectedDbPath, dbPath)
}
