package report

import (
	"encoding/csv"
	"os"

	"github.com/luistm/banksaurus/next/adapter/CGDcsv"
	"github.com/luistm/banksaurus/next/adapter/transactionpresenter"
	"github.com/luistm/banksaurus/next/listtransactions"
	"github.com/luistm/banksaurus/next/reportgrouped"
)

// Command handles reports
type Command struct{}

// Execute the report command
func (rc *Command) Execute(arguments map[string]interface{}) error {
	var grouped bool

	if arguments["--grouped"].(bool) {
		grouped = true
	}

	filePath := arguments["<file>"].(string)
	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}

	cgdCSVRepository, err := CGDcsv.New(lines)
	if err != nil {
		return err
	}

	p, err := transactionpresenter.NewPresenter()
	if err != nil {
		return err
	}

	if grouped {
		i, err := reportgrouped.NewInteractor(cgdCSVRepository, p)
		if err != nil {
			return err
		}

		err = i.Execute()
		if err != nil {
			return err
		}

	} else {
		i, err := listtransactions.NewInteractor(p, cgdCSVRepository)
		if err != nil {
			return err
		}

		err = i.Execute()
		if err != nil {
			return err
		}
	}

	vm, err := p.ViewModel()
	if err != nil {
		return err
	}

	vm.Write(os.Stdout)

	return nil
}
