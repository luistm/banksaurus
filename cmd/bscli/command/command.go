package command

import (
	"errors"

	"github.com/luistm/banksaurus/cmd/bscli/command/load"
	"github.com/luistm/banksaurus/cmd/bscli/command/report"
	"github.com/luistm/banksaurus/cmd/bscli/command/seller"
)

var errCommandNotFound = errors.New("command not found")
var errCommandIsUndefined = errors.New("command is undefined")

// CLIRequest is the interface to pass information to a command execution
type CLIRequest []string

// Commander executes a request from the command line
type Commander interface {
	Execute(map[string]interface{}) error
}

// New creates a new command handler
func New(cliRequest CLIRequest) (Commander, error) {

	if len(cliRequest) == 0 {
		return nil, errCommandIsUndefined
	}

	command := cliRequest[0]
	switch command {
	case "report":
		return &report.Command{}, nil
	case "load":
		return &load.Command{}, nil
	case "seller":
		return &seller.Command{}, nil
	default:
		return nil, errCommandNotFound
	}
}
