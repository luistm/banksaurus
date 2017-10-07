package reports

import (
	"go-bank-cli/infrastructure"
)

// ParseAccountMovements imports data from a data source
func ParseAccountMovements(filePath string) ([][]string, error) {

	fileRecords, err := infrastructure.OpenFile(filePath)
	if err != nil {
		return [][]string{}, err
	}

	// TODO: Transfor records into transactions

	return fileRecords, nil
}
