package main

import (
	"fmt"
	"log"
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

var version = "bankcli 1.1.0"

func configureLog() {
	f, err := os.OpenFile("filename", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// TODO: What to do here? I don't want to Fatal
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
}

// func setup(){
// 	// this is the only place where i do this!!
// 	if dev
// 	this
// 	else
// 	that
// }

func main() {

	arguments, err := docopt.Parse(intro+usage+options, nil, true, version, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	// TODO: Make the setup of the application here. What happens if this is first execution
	// setup()
	//       When running tests, does this create the log files ?
	//       Maybe in dev i should print the logs and not send it to the file.
	configureLog()

	command, _ := commands.New(os.Args[1:])
	response := command.Execute(arguments)
	if response.String() != "" {
		fmt.Println(response.String())
		os.Exit(2)
	}
}
