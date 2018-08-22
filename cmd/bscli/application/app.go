package application

import (
	"os"
)

const (
	// Version of the application
	Version = "bscli 2.0.0"
)

// New creates an instance of App.
func New() (*App, error) {

	application := &App{}

	err := application.createPathAndFolders()
	if err != nil {
		return &App{}, err
	}

	db, err := Database()
	if err != nil {
		return &App{}, err
	}

	err = buildSchema(db)
	if err != nil {
		return &App{}, err
	}

	return application, nil
}

// App represents an application
type App struct {
	ProjectPath string
}

// createPathAndFolders initializes the application.
func (a *App) createPathAndFolders() error {

	// Create home dir if not exists
	_, err := os.Stat(Path())

	// Create Path if it does not exist
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(Path(), 0700)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	return nil
}
