package commands

import "errors"

var errCommandNotFound = errors.New("command not found")
var errCommandIsUndefined = errors.New("command is undefined")

// CLIRequest is the interface to pass information to a command execution
type CLIRequest []string // TODO: This interface is not very useful. Think about is!

// CommandHandler executes a request from the command line
type CommandHandler interface {
	Execute(map[string]interface{}) error
}

// New creates a new command handler
func New(cliRequest CLIRequest) (CommandHandler, error) {

	if len(cliRequest) == 0 {
		return nil, errCommandIsUndefined
	}

	command := cliRequest[0]
	switch command {
	case "reportgrouped":
		return &Report{}, nil
	case "loaddata":
		return &Load{}, nil
	case "seller":
		return &Seller{}, nil
	case "transaction":
		return &TransactionCommand{}, nil
	default:
		return nil, errCommandNotFound
	}
}
