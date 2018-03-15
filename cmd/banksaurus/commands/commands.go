package commands

import "errors"

var errCommandNotFound = errors.New("command not found")
var errCommandIsUndefined = errors.New("command is undefined")

type cliRequest []string

// CommandHandler executes a request from the command line
type CommandHandler interface {
	Execute(map[string]interface{}) error
}

// New creates a new command handler
func New(cliRequest cliRequest) (CommandHandler, error) {

	if len(cliRequest) == 0 {
		return nil, errCommandIsUndefined
	}

	command := cliRequest[0]
	switch command {
	case "report":
		return &Report{}, nil
	case "load":
		return &Load{}, nil
	case "seller":
		return &Seller{}, nil
	default:
		return nil, errCommandNotFound
	}
}
