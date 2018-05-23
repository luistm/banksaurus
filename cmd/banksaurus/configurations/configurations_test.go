package configurations

import (
	"os"
	"os/user"
	"path"
	"testing"

	"github.com/luistm/testkit"
)

func TestUnitGetDataBasePath(t *testing.T) {

	usr, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	expectedDbName := "services"
	expectedDbPath := path.Join(usr.HomeDir, ".services")

	dbName, dbPath := DatabasePath()

	testkit.AssertEqual(t, expectedDbName, dbName)
	testkit.AssertEqual(t, expectedDbPath, dbPath)

	os.Setenv("GO_BANK_CLI_DEV", "true")
	defer os.Setenv("GO_BANK_CLI_DEV", "")

	expectedDbName = "services"
	expectedDbPath = os.TempDir()

	dbName, dbPath = DatabasePath()

	testkit.AssertEqual(t, expectedDbName, dbName)
	testkit.AssertEqual(t, expectedDbPath, dbPath)
}
