package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/luistm/banksaurus/cmd/banksaurus/command"
	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
	"github.com/luistm/banksaurus/app"
)

var intro = "    \n    Your command line finance manager.\n\n"

var usage = `Usage:
	banksaurus -h | --help
	banksaurus report --input <file> [ --grouped ]
	banksaurus load --input <file>
	banksaurus seller change <id> --pretty <name>
	banksaurus seller new <name>
	banksaurus seller show
	banksaurus transaction show`

var options = `

Options:
	--grouped     The report is grouped by seller.
	--input       The path to the records list.
	--name        Specifies the name.
	-h --help     Show this screen.`

func setup() error {
	if configurations.IsDev() {
		return nil
	}

	// Create home dir if not exists
	_, err := os.Stat(configurations.ApplicationHomePath())
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(configurations.ApplicationHomePath(), 0700)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

var errGeneric = errors.New("error while performing operation")

func main() {
	err := setup()
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to setup application")
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
