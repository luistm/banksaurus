package app

import (
	"os"

	"github.com/luistm/banksaurus/configurations"
)

const (
	// Version of the application
	Version = "banksaurus 1.2.0"
)

// New creates an instance of App in the specified path.
// It should receive the path to the configuration file.
func New(configuration string) (*App, error) {
	// TODO: Validate the path to the configuration file.
	application :=  &App{}
	err := application.Init()
	if err != nil {
		return &App{}, err
	}

	return application, nil
}

// App represents an application
type App struct{}

// Init initializes the application
func (a *App) Init() error {

	if configurations.IsDev() {
		return nil
	}

	// Create home dir if not exists
	_, err := os.Stat(configurations.ApplicationHomePath())
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(configurations.ApplicationHomePath(), 0700)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
