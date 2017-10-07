package reports

import (
	"fmt"
	"go-bank-cli/lib/transactions"
	"io"

	"github.com/shopspring/decimal"
)

// IRepository is the interface for report repositories
type IRepository interface {
	AllTransactions() ([]*Transaction, error)
}

// Interactor ...
type Interactor struct {
	Repository IRepository
}

// MonthlyReport produces a report for the current month
func (i *Interactor) MonthlyReport() (*Report, error) {

	// Import transactions from this month
	_, err := i.Repository.AllTransactions()
	if err != nil {
		return &Report{}, fmt.Errorf("Failed to create report: %s", err)
	}
	// For each expense transaction
	// If transaction isDebt: Report.AddExpense(t)
	// If transaction isCredit: Report.AddCredit(t)
	//
	// Return report, nil

	return &Report{}, nil
}

// LoadReport loads an external list of an account movement
func LoadReport(file io.Reader) error {

	records, err := ImportData(file)
	if err != nil {
		return fmt.Errorf("Failed to import data: %s", err)
	}

	var report map[string]decimal.Decimal
	report = make(map[string]decimal.Decimal)
	var credit decimal.Decimal
	var expense decimal.Decimal

	// Read all transactions
	for lineCount, record := range records {

		r := transactions.Record{Record: record}

		if !r.Valid() || lineCount < 4 {
			continue
		}

		t := transactions.Transaction{}
		transaction := t.New(r)

		if transaction.IsFromThisMonth() {
			report[transaction.Description] = report[transaction.Description].Add(transaction.Value())
			if transaction.IsDebt() {
				expense = expense.Add(transaction.Value())
			} else {
				credit = credit.Add(transaction.Value())
			}
		}

	}

	for transactionDescription, transactionValue := range report {
		fmt.Printf("%24s %8s \n", transactionDescription, transactionValue.String())
	}

	fmt.Println("Expense is ", expense.String())
	fmt.Println("Credit is ", credit.String())

	return nil
}
