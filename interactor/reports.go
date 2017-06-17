package interactor

import (
	"expensetracker/entities"
	"fmt"
)

// MonthlyReport builds a sum of expenses and credits
func MonthlyReport(records [][]string) error {

	var report map[string]float64
	report = make(map[string]float64)
	var credit float64
	var expense float64

	// Read all transactions
	for lineCount, record := range records {

		r := entities.Record{Record: record}

		if len(r.Record) != 8 || lineCount < 4 {
			continue
		}

		t := entities.Transaction{}
		transaction := t.New(r)

		report[transaction.Description] += transaction.Value()
		if transaction.TransactionType == entities.DEBT {
			expense += transaction.Value()
		} else {
			credit += transaction.Value()
		}
	}

	for transactionDescription, transactionValue := range report {
		fmt.Printf("%24s %8.2f \n", transactionDescription, transactionValue)
	}

	fmt.Println("Expense is ", expense)
	fmt.Println("Credit is ", credit)

	return nil
}
