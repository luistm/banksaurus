package app

import (
	"os"

	"strings"

	"path"
	"path/filepath"

	"github.com/luistm/banksaurus/configurations"
)

const (
	// Version of the application
	Version = "banksaurus 1.2.0"
)

// New creates an instance of App in the specified path.
// It should receive the full path to the configuration file.
func New(configurationPath string) (*App, error) {

	err := ValidatePath(configurationPath)
	if err != nil {
		return &App{}, err
	}

	projectPath, err := buildProjectPath(filepath.Dir(configurationPath))
	if err != nil {
		return &App{}, err
	}

	application := &App{projectPath}
	err = application.Init()
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
