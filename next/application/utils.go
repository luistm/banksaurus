package application

import (
	"os"
	"os/user"
	"path"
)

// ConfigHome defines the base directory relative to which
// user specific configuration files should be stored.
func ConfigHome() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return path.Join(usr.HomeDir, ".bank")
}

// DatabasePath returns the path and name for the database
// taking into account the type of environment
func DatabasePath() (string, string) {
	dbName := "bank"

	// TODO: Move the database name to the configurations file
	if isDev() {
		// TODO: When testing, use an in memory database.
		return dbName, os.TempDir()
	}

	return dbName, ConfigHome()
}

// isDev ...
func isDev() bool {
	// TODO: Find a solution to delete this function
	if os.Getenv("BANKSAURUS_ENV") == "dev" {
		return true
	}
	return false
}
