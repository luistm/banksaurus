package account

import "errors"

// ErrUnrecognizedCommand ...
var ErrUnrecognizedCommand = errors.New("unrecognized command")

const (
	list   = "list"
	create = "create"
)

// NewCommand creates a new instance of the seller command
func NewCommand(subCommand string) (Commander, error) {
	if subCommand != list && subCommand != create {
		return &CreateAccountCommand{}, ErrUnrecognizedCommand
	}

	var command Commander
	if subCommand == create {
		command = &CreateAccountCommand{}
	}

	if subCommand == list {
		command = &ListAccountsCommand{}
	}

	return command, nil
}
