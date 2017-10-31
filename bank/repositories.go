package bank

// CSVHandler to handle csv files
type CSVHandler interface {
	Lines() ([][]string, error)
}
