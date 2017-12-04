package main

// CommandHandler executes a request from the command line
type CommandHandler interface {
	Execute(*Request) error
}

// Request is a set of options and values to a command
type Request []string

// ReportCommand handles reports
type ReportCommand struct{}

// Execute a report command
func (rc *ReportCommand) Execute(r *Request) error {
	return nil
}

// NewHandler creates a new CommandHandler
func NewHandler(command string) (CommandHandler, error) {

	switch command {
	case "report":
		return &ReportCommand{}, nil
	}

	return nil, nil
}
