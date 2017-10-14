package main

// This software is an expense tracker i made to read the transactions exported from
// my bank account.

import (
	"fmt"
	"os"

	docopt "github.com/docopt/docopt-go"
	"github.com/luistm/go-bank-cli/infrastructure"
)

func errorf(format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, format, args)
	os.Exit(2)
}

var intro = "Your command line finance manager.\n"

var usage = `Usage:
	bank -h | --help
	bank report --input <file>
	bank category new <name>
	bank category show
	bank description show`

var options = `
Options:
	--input       The path to the transactions list.
	-h --help     Show this screen.`

func main() {

	arguments, _ := docopt.Parse(intro+usage+options, nil, true, "Go CLI Bank 0.0.1", false)

	err := infrastructure.InitStorage(DatabaseName, DatabasePath)
	if err != nil {
		errorf("Error:", err)
	}

	out := ""
	if arguments["category"].(bool) && arguments["new"].(bool) {
		out, err = createCategoryHandler(arguments["<name>"].(string))
	}

	if arguments["category"].(bool) && arguments["show"].(bool) {
		out, err = showCategoryHandler()
	}

	if arguments["report"].(bool) {
		out, err = showReportHandler(arguments["<file>"].(string))
	}

	if err != nil {
		errorf("Error:", err)
	}
	fmt.Printf(out)
}
