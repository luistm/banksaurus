package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/luistm/banksaurus/cmd/banksaurus/command"
	app "github.com/luistm/banksaurus/next/application"
)

var intro = "    \n    Your command line finance manager.\n\n"

var usage = `Usage:
	banksaurus -h | --help
	banksaurus report --input <file> [ --grouped ]
	banksaurus load --input <file>
	banksaurus seller change <id> --pretty <name>
	banksaurus seller new <name>
	banksaurus seller show`

var options = `

Options:
	--grouped     The report is grouped by seller.
	--input       The path to the records list.
	--name        Specifies the name.
	-h --help     Show this screen.`

var errGeneric = errors.New("error while performing operation")

func main() {
	_, err := app.New("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Setup application: %s\n", err.Error())
		os.Exit(2)
	}

	// TODO: Inject dependencies here
	// err := application.Add(aConstructor, "constructor.slug")
	// defer application.Close()

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
