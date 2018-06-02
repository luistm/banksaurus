package app

import (
	"os"

	"strings"

	"path"
	"path/filepath"

	"os/user"
)

const (
	// Version of the application
	Version = "banksaurus 1.2.0" // TODO: define this automagically on make build
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
	err := application.Init()
	if err != nil {
		return &App{}, err
	}

	return application, nil
}

// A project must have in the project path
// - /configurations
// - /cmd
// - /infrastructure
func buildProjectPath(pth string) (string, error) {

	p := path.Join(pth, "..")

	err := ValidatePath(path.Join(p, "/configurations"))
	if err != nil {
		return "", ErrInvalidAppStructure
	}

	err = ValidatePath(path.Join(p, "/cmd"))
	if err != nil {
		return "", ErrInvalidAppStructure
	}

	err = ValidatePath(path.Join(p, "/infrastructure"))
	if err != nil {
		return "", ErrInvalidAppStructure
	}

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

// Init initializes the application
func (a *App) Init() error {

	if IsDev() {
		return nil
	}

	// Create home dir if not exists
	_, err := os.Stat(ConfigHome())
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(ConfigHome(), 0700)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// TODO: Must create the configuration file here
	// touch applicationpath/config.json

	return nil
}

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
	if IsDev() {
		return dbName, os.TempDir()
	}

	return dbName, ConfigHome()
}

// IsDev returns if in dev environment or not
func IsDev() bool {
	if os.Getenv("BANKSAURUS_ENV") == "dev" {
		return true
	}
	return false
}
