package configurations

import (
	"os"
	"os/user"
	"path"
)

func isDev() bool {
	if os.Getenv("GO_BANK_CLI_DEV") == "true" {
		return true
	}
	return false
}

// DatabasePath returns the path nad name for the database
// taking into account the type of environment
func DatabasePath() (string, string) {
	if isDev() {
		return "bank", "/tmp"
	}

	usr, err := user.Current()
	if err != nil {
		// TODO: no panic here...
		panic(err)
	}
	dbName := "bank"
	dbPath := path.Join(usr.HomeDir, ".bank")

	return dbName, dbPath
}
