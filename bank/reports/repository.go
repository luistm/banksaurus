package reports

// CSVHandler to handle csv files
type CSVHandler interface {
	GetAll() ([][]string, error)
}

type repository struct {
	storage CSVHandler
}

// ParseAccountMovements imports data from a data source
func (r *repository) GetAll() error {

	// fileRecords, err := r.storage.GetAll()
	// if err != nil {
	// 	return err
	// }

	// var report map[string]decimal.Decimal
	// report = make(map[string]decimal.Decimal)
	// var credit decimal.Decimal
	// var expense decimal.Decimal

	// // Read all transactions
	// for lineCount, record := range fileRecords {

	// 	r := transactions.Record{Record: record}

	// 	if !r.Valid() || lineCount < 4 {
	// 		continue
	// 	}

	// 	t := transactions.Transaction{}
	// 	transaction := t.New(r)

	// 	// descriptions.New(transaction.Description)

	// 	if transaction.IsFromThisMonth() {
	// 		report[transaction.Description] = report[transaction.Description].Create(transaction.Value())
	// 		if transaction.IsDebt() {
	// 			expense = expense.Create(transaction.Value())
	// 		} else {
	// 			credit = credit.Create(transaction.Value())
	// 		}
	// 	}

	// }

	// for transactionDescription, transactionValue := range report {
	// 	fmt.Printf("%24s %8s \n", transactionDescription, transactionValue.String())
	// }

	// fmt.Println("Expense is ", expense.String())
	// fmt.Println("Credit is ", credit.String())

	return nil

}
