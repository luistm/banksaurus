package transactions

import (
	"github.com/luistm/go-bank-cli/bank"
	"github.com/luistm/go-bank-cli/lib/customerrors"
	"github.com/luistm/go-bank-cli/lib/sellers"
)

type repository struct {
	storage      bank.CSVHandler
	transactions []*Transaction
}

func (r *repository) GetAll() ([]*Transaction, error) {

	if r.storage == nil {
		return []*Transaction{}, customerrors.ErrInfrastructureUndefined
	}

	_, err := r.storage.Lines()
	if err != nil {
		return []*Transaction{}, &customerrors.ErrInfrastructure{Msg: err.Error()}
	}

	// TODO: Validate if Lines() output is the expected one
	// r.validateLines(lines)
	// if err != nil{}
	// TODO: r.buildTransactions(lines)
	// if err != nil{}

	return []*Transaction{}, nil
}

func (r *repository) buildTransactions(lines [][]string) error {
	for _, l := range lines {
		t := &Transaction{
			s: sellers.New(l[2], l[2]),
		}
		r.transactions = append(r.transactions, t)
	}

	return nil
}
