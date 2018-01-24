package configurations

import (
	"os"
	"os/user"
	"path"
	"testing"

	"github.com/luistm/go-bank-cli/elib/testkit"
)

func TestUnitGetDataBasePath(t *testing.T) {

	usr, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	expectedDbName := "bank"
	expectedDbPath := path.Join(usr.HomeDir, ".bank")

	dbName, dbPath := DatabasePath()

	testkit.AssertEqual(t, expectedDbName, dbName)
	testkit.AssertEqual(t, expectedDbPath, dbPath)

	os.Setenv("GO_BANK_CLI_DEV", "true")
	defer os.Setenv("GO_BANK_CLI_DEV", "")

	expectedDbName = "bank"
	expectedDbPath = "/tmp"

	dbName, dbPath = DatabasePath()

	testkit.AssertEqual(t, expectedDbName, dbName)
	testkit.AssertEqual(t, expectedDbPath, dbPath)
}
