package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
	app "github.com/luistm/banksaurus/cmd/bscli/application"
	"github.com/luistm/banksaurus/cmd/bscli/command"
)

var intro = "    \n    Your command line finance manager.\n\n"

var usage = `Usage:
	bscli -h | --help
	bscli report
	bscli report --input <file> [ --grouped ]
	bscli load --input <file>
	bscli seller change <id> --pretty <name>
	bscli seller new <name>
	bscli seller show`

var options = `

Options:
	--grouped     The report is grouped by seller.
	--input       The path to the records list.
	--name        Specifies the name.
	-h --help     Show this screen.`

var errGeneric = errors.New("error while performing operation")

func main() {
	_, err := app.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Setup application: %s\n", err.Error())
		os.Exit(2)
	}

	arguments, err := docopt.Parse(intro+usage+options, nil, true, app.Version, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	cmd, err := command.New(os.Args[1:])
	if err != nil {

		fmt.Fprintf(os.Stderr, errGeneric.Error())
		os.Exit(2)
	}

	err = cmd.Execute(arguments)
	if err != nil {
		fmt.Fprintf(os.Stderr, errGeneric.Error()+"\n")
		os.Exit(2)
	}
}
