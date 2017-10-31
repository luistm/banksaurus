package reports

import (
	"errors"

	"github.com/luistm/go-bank-cli/bank"
)

// NewInteractor creates an interactor for reports
func NewInteractor(storage bank.CSVHandler) *interactor {
	r := &repository{storage: storage}
	return &interactor{repository: r}
}

type interactor struct {
	repository *repository
}

// CurrentMonth produces a report for the current month
func (i *interactor) CurrentMonth() (*Report, error) {

	if i.repository == nil {
		return &Report{}, errors.New("repository is not defined")
	}

	// Import transactions from this month
	// _, err := i.repository.GetAll()
	// if err != nil {
	// 	return &Report{}, fmt.Errorf("Failed to create report: %s", err)
	// }
	// For each expense transaction
	// If transaction isDebt: Report.AddExpense(t)
	// If transaction isCredit: Report.AddCredit(t)
	//
	// Return report, nil

	return &Report{}, nil
}
