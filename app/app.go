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
	Version = "banksaurus 1.2.0"
)

// New creates an instance of App in the specified path.
// It should receive the full path to the configuration file.
// If the full path to the configuration path is not received,
// defaults will be assumed.
func New(configurationPath string) (*App, error) {

	var projectPath string


	// Environment variable has priority above others
	configPath := os.Getenv("BANKSAURUS_CONFIG")
	if configPath != ""{
		configurationPath = configPath
	}

	// Assume defaults
	if configurationPath == ""{
		projectPath = path.Join(ApplicationHomePath(), "config.json")
	}

	if configurationPath != "" {
		err := ValidatePath(configurationPath)
		if err != nil {
			return &App{}, err
		}

		projectPath, err = buildProjectPath(filepath.Dir(configurationPath))
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
// A path is considered valid if:
// a) the path exists
// b) it's not a file
// If a path is valid the return is nil
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

	//if !fileInfo.IsDir() {
	//	return ErrPathIsFile
	//}

	return nil
}

// Init initializes the application
func (a *App) Init() error {

	if IsDev() {
		return nil
	}

	// Create home dir if not exists
	_, err := os.Stat(ApplicationHomePath())
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(ApplicationHomePath(), 0700)
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

// ApplicationHomePath builds the path to application data in the user home,
// something like ~/.services
func ApplicationHomePath() string {
	usr, err := user.Current()
	if err != nil {
		// TODO: no panic here...
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

	return dbName, ApplicationHomePath()
}

// IsDev returns if in dev environment or not
func IsDev() bool {
	if os.Getenv("BANKSAURUS_ENV") == "dev" {
		return true
	}
	return false
}
