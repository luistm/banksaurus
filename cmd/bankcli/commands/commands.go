package commands

import "errors"

var errCommandNotFound = errors.New("command not found")
var errCommandIsUndefined = errors.New("command is undefined")

type cliRequest []string

//Response has the result of a command execution
type Response struct {
	err    error
	output string
}

func (res *Response) String() string {
	if res.err != nil {
		return res.err.Error()
	}

	return res.output

}

// CommandHandler executes a request from the command line
type CommandHandler interface {
	Execute(map[string]interface{}) *Response
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
	case "category":
		return &Category{}, nil
	default:
		return nil, errCommandNotFound
	}
}
