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
	dbName := "bank"
	if isDev() {
		return dbName, os.TempDir()
	}

	return dbName, ApplicationHomePath()
}

// LogPath returns the path to the log file
func LogPath() string {
	return path.Join(ApplicationHomePath(), ".log")
}

// ApplicationHomePath builds the path to application data in the user home,
// something like ~/.bank
func ApplicationHomePath() string {
	usr, err := user.Current()
	if err != nil {
		// TODO: no panic here...
		panic(err)
	}
	return path.Join(usr.HomeDir, ".bank")
}
