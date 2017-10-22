package reports

import (
	"errors"
	"fmt"
)

// IRepository is the interface for report repositories
type IRepository interface {
	AllTransactions() ([]*Transaction, error)
}

// Interactor ...
type Interactor struct {
	repository IRepository
}

// MonthlyReport produces a report for the current month
func (i *Interactor) MonthlyReport() (*Report, error) {

	if i.repository == nil {
		return &Report{}, errors.New("repository is not defined")
	}

	// Import transactions from this month
	_, err := i.repository.AllTransactions()
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
func LoadReport(filePath string) error {

	err := ParseAccountMovements(filePath)
	if err != nil {
		return fmt.Errorf("Failed to import data: %s", err)
	}

	return nil
}
