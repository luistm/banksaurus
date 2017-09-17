// Package cats provides a framework for testing command line applications
package cats

import "errors"

// New creates a new application given the executable name
func New(command string) *Application {
	return &Application{command: command}
}

// Application represents an executable to be tested
type Application struct {
	command string
}

// Run executes the application with arguments
func (app *Application) Run(args ...string) error {

	// Capture STDERR
	// Capture STDOUT

	return errors.New("Error")
}

//assert.Contains(t, expected, output)
//assert.NoError(t, err)
//assert.Equal(t, expected, output)
