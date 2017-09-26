package main

// This software is an expense tracker i made to read the transactions exported from
// my bank account.

import (
	"fmt"
	"go-bank-cli/infrastructure"
	"os"

	docopt "github.com/docopt/docopt-go"
)

func errorf(format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, format, args)
	os.Exit(2)
}

var intro = "Your command line money manager.\n"

var usage = `Usage:
	bank -h | --help
	bank report --input <file>
	bank category new <name>`

var options = `
Options:
	--input    The path to the transactions list.
	-h --help     Show this screen.`

func main() {

	arguments, _ := docopt.Parse(intro+usage+options, nil, true, "Go CLI Bank 0.0.1", false)

	// TODO: Get database values from a configuration
	err := infrastructure.InitStorage(DatabaseName, DatabasePath)
	if err != nil {
		errorf("Error:", err)
	}

	out := ""
	if arguments["category"].(bool) && arguments["new"].(bool) {
		out, err = createCategoryHandler(arguments["<name>"].(string))
	}

	if arguments["report"].(bool) {
		out, err = showReportHandler(arguments["<file>"].(string))
	}

	if err != nil {
		errorf("Error:", err)
	}
	fmt.Printf(out)
}
