package account

import "errors"

// ErrUnrecognizedCommand ...
var ErrUnrecognizedCommand = errors.New("unrecognized command")

const (
	show   = "show"
	create = "create"
)

// NewCommand creates a new instance of the seller command
func NewCommand(subCommand string) (Commander, error) {
	if subCommand != show && subCommand != create {
		return &CreateAccountCommand{}, ErrUnrecognizedCommand
	}

	var command Commander
	if subCommand == create {
		command = &CreateAccountCommand{}
	}

	if subCommand == show {
		command = &ListAccountsCommand{}
	}

	return command, nil
}
