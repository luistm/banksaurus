package application

import (
	"os"
	"os/user"
	"path"
)

const applicationHomePath = ".bank"

var appPath = ""

// Path defines the base directory relative to which
// user specific configuration files should be stored.
func Path() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	if isDev() {
		appPath = path.Join(usr.HomeDir, os.TempDir())
	} else {
		appPath = path.Join(usr.HomeDir, applicationHomePath)
	}

	return appPath
}

func isDev() bool {
	// TODO: Find a solution to delete this function
	if os.Getenv("BANKSAURUS_ENV") == "dev" {
		return true
	}
	return false
}
