package main

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
	execute(map[string]interface{}) *Response
}

func newCommand(cliRequest cliRequest) (CommandHandler, error) {

	if len(cliRequest) == 0 {
		return &ReportCommand{}, errCommandIsUndefined
	}

	command := cliRequest[0]
	switch command {
	case "report":
		return &ReportCommand{commandType: command}, nil
	default:
		return &ReportCommand{}, errCommandNotFound
	}
}
