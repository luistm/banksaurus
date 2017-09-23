package main

// This software is an expense tracker i made to read the transactions exported from
// my bank account.

import (
	"fmt"
	"go-cli-bank/commands"
	"os"

	docopt "github.com/docopt/docopt-go"
)

func errorf(format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, format, args)
	os.Exit(2)
}

var intro = "Your command line money manager.\n"

var usage = `Usage:
	go-cli-bank -h | --help
	go-cli-bank report --input <file>
	go-cli-bank category new <name>`

var options = `
Options:
	--input    The path to the transactions list.
	-h --help     Show this screen.`

func main() {

	arguments, _ := docopt.Parse(intro+usage+options, nil, true, "Go CLI Bank 0.0.1", false)

	out := ""
	var err error
	if arguments["category"].(bool) && arguments["new"].(bool) {
		out, err = commands.CreateCategoryHandler(arguments["<name>"].(string))
	}

	if arguments["report"].(bool) {
		out, err = commands.ShowReportHandler(arguments["<file>"].(string))
	}

	if err != nil {
		errorf("Error:", err)
	}
	fmt.Printf(out)
}
