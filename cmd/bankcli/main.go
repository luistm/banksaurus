package main

import (
	"fmt"
	"os"

	docopt "github.com/docopt/docopt-go"
	"github.com/luistm/go-bank-cli/cmd/bankcli/commands"
)

var intro = "    \n    Your command line finance manager.\n\n"

var usage = `Usage:
	bankcli -h | --help
	bankcli report --input <file> [ --grouped ]
	bankcli load --input <file>
	bankcli seller change <id> --pretty <name>
	bankcli seller new <name>
	bankcli seller show`

var options = `

Options:
	--grouped     The report is present grouped by seller
	--input       The path to the records list.
	--name        Specifies the name.
	-h --help     Show this screen.`

var version = "1.0.0"

func main() {
	var out string
	var err error

	arguments, err := docopt.Parse(intro+usage+options, nil, true, version, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	command, err := commands.New(os.Args[1:])
	if err == nil {
		out = command.Execute(arguments).String()
	}
	fmt.Printf(out)
}
