package reports

import (
	"encoding/csv"
	"io"
)

// ImportData imports data from a data source
func ImportData(r io.Reader) ([][]string, error) {

	reader := csv.NewReader(r)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
