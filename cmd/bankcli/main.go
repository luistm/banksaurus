package main

// This software is an expense tracker i made to read the transactions exported from
// my bank account.

import (
	"fmt"
	"os"

	docopt "github.com/docopt/docopt-go"
	"github.com/luistm/go-bank-cli/cmd/bankcli/commands"
)

func errorf(format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, format, args)
	os.Exit(2)
}

var intro = "Your command line finance manager.\n"

var usage = `Usage:
	bankcli -h | --help
	bankcli report --input <file> [ --grouped ]
	bankcli load --input <file>
	bankcli category new <name>
	bankcli category show
	bankcli seller change <id> --pretty <name>
	bankcli seller new <name>
	bankcli seller show`

var options = `
Options:
	--grouped     The report is present grouped by seller
	--input       The path to the records list.
	--name        Specifies the name.
	-h --help     Show this screen.`

func main() {
	var out string
	var err error

	arguments, _ := docopt.Parse(intro+usage+options, nil, true, "Go CLI Bank 0.0.1", false)

	if arguments["seller"].(bool) && arguments["new"].(bool) {
		out, err = commands.CreateSellerHandler(arguments["<name>"].(string))
	}

	if arguments["seller"].(bool) && arguments["show"].(bool) {
		err = commands.ShowSellersHandler()
		// TODO: Remove this workaround after moving this a command
		if err != nil {
			errorf("Error:", err)
		}
		os.Exit(0)
	}

	if arguments["seller"].(bool) && arguments["change"].(bool) {
		out, err = commands.SellerChangePrettyName(
			arguments["<id>"].(string),
			arguments["<name>"].(string),
		)
	}

	if err != nil {
		errorf("Error:", err)
	}

	command, err := commands.New(os.Args[1:])
	if err == nil {
		out = command.Execute(arguments).String()
	}
	fmt.Printf(out)
}
