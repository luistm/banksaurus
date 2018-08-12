package application

import (
	"os"

	"strings"

	"errors"
	"path"
	"path/filepath"
)

const (
	// Version of the application
	Version = "banksaurus 1.2.0"

	// TODO: Application version should be defined auto magically when a release is made
)

var (
	// ErrMalformedPath ...
	ErrMalformedPath = errors.New("path is malformed")

	// ErrPathDoesNotExist ...
	ErrPathDoesNotExist = errors.New("path does not exist")

	// ErrInvalidAppStructure ...
	ErrInvalidAppStructure = errors.New("invalid application structure")
)

// New creates an instance of App.
//
// An app needs a configuration file.
// The path to the configuration file can be defined by:
// 1) passing by argument to the new function;
// 2) Defining the BANKSAURUS_CONFIG environment variable
// Both ways should receive the full path to the configuration file.
//
// If a configuration file path is not defined,
// the default value is $HOME/.bank/config.json
//
func New(configurationFilePath string) (*App, error) {

	var projectPath string

	// Environment variable has priority above others
	configPath := os.Getenv("BANKSAURUS_CONFIG")
	if configPath != "" {
		configurationFilePath = configPath
	}

	// Assume defaults
	if configurationFilePath == "" {
		projectPath = path.Join(ConfigHome(), "config.json")
	}

	if configurationFilePath != "" {
		err := ValidatePath(configurationFilePath)
		if err != nil {
			return &App{}, err
		}

		projectPath, err = buildProjectPath(filepath.Dir(configurationFilePath))
		if err != nil {
			return &App{}, err
		}
	}

	application := &App{projectPath}
	if isDev() {
		dbName, dbPath := DatabasePath()
		_, err := NewSchema(dbPath, dbName, false)
		if err != nil {
			return &App{}, err
		}
		return application, nil
	}

	err := application.Init()
	if err != nil {
		return &App{}, err
	}

	return application, nil
}

// A project must have in the project path:
// a) /configurations
// b) /cmd
// c) /infrastructure
func buildProjectPath(pth string) (string, error) {

	p := path.Join(pth, "..")

	err := ValidatePath(path.Join(p, "/configurations"))
	if err != nil {
		return "", ErrInvalidAppStructure
	}

	//err = ValidatePath(path.Join(p, "/cmd"))
	//if err != nil {
	//	return "", ErrInvalidAppStructure
	//}
	//
	//err = ValidatePath(path.Join(p, "/infrastructure"))
	//if err != nil {
	//	return "", ErrInvalidAppStructure
	//}

	return p, nil
}

// App represents an application
type App struct {
	ProjectPath string
}

// ValidatePath checks if a path is valid
// A path is considered valid if exists.
func ValidatePath(path string) error {

	if !strings.HasPrefix(path, "/") {
		return ErrMalformedPath
	}

	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return ErrPathDoesNotExist
	}
	if err != nil {
		return err
	}

	return nil
}

// Init initializes the application.
func (a *App) Init() error {

	// Create home dir if not exists
	_, err := os.Stat(ConfigHome())

	// Create ConfigHome if it does not exist
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir(ConfigHome(), 0700)
		if err != nil {
			return err
		}
	}

	// Handle all other errors
	if err != nil {
		return err
	}

	// TODO: Create configuration files
	//       ~/.bank/config.json
	dbName, dbPath := DatabasePath()
	_, err = NewSchema(dbPath, dbName, false)
	if err != nil {
		return err
	}

	return nil
}
