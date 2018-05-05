package transaction

import (
	"errors"
	"fmt"
	"strings"

	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/banklib/moneyamount"
	"github.com/luistm/banksaurus/banklib/seller"
)

// NewRepository creates an instance of the Transactions repository
func NewRepository(storage banklib.CSVHandler, relationalStorage banklib.SQLInfrastructer) *Transactions {
	return &Transactions{storage: storage, relationalStorage: relationalStorage}
}

// Transactions repository
type Transactions struct {
	storage           banklib.CSVHandler
	Transactions      []banklib.Entity // TODO: Make private
	relationalStorage banklib.SQLInfrastructer
}

// Save to save a transaction
func (r *Transactions) Save(t banklib.Entity) error {

	return errors.New("save not implemented")
}

// Get to fetch a single transaction
func (r *Transactions) Get(s string) (banklib.Entity, error) {
	// TODO: Implement this
	return &Transaction{}, errors.New("get not implemented")
}

// GetAll to fetch all Transactions
func (r *Transactions) GetAll() ([]banklib.Entity, error) {
	// TODO: Should return an iterator

	if r.storage == nil {
		return []banklib.Entity{}, banklib.ErrInfrastructureUndefined
	}

	lines, err := r.storage.Lines()
	if err != nil {
		return []banklib.Entity{}, &banklib.ErrInfrastructure{Msg: err.Error()}
	}

	// TODO: Validate if Lines() output is the expected one
	if len(lines) == 0 {
		return []banklib.Entity{}, &banklib.ErrInfrastructure{Msg: "empty data to parse"}
	}
	if len(lines) < 6 {
		return []banklib.Entity{}, &banklib.ErrInfrastructure{Msg: "data has an unknown format"}
	}

	err = r.BuildTransactions(lines[5 : len(lines)-2])
	if err != nil {
		return []banklib.Entity{}, &banklib.ErrInfrastructure{Msg: err.Error()}
	}

	return r.Transactions, nil
}

// BuildTransactions creates Transactions from an array of records
func (r *Transactions) BuildTransactions(lines [][]string) error {
	// TODO: This is to be private
	for i, line := range lines {

		// TODO: Handle credit
		if line[3] == "" {
			continue
		}

		value, err := moneyamount.New(line[3])
		if err != nil {
			return fmt.Errorf("failed to create decimal from string: %s", err.Error())
		}

		slug := strings.TrimSuffix(line[2], " ")
		t := &Transaction{
			id:     uint64(i),
			value:  value.AmountToDecimal(),
			Seller: seller.New(slug, slug),
		}
		r.Transactions = append(r.Transactions, t)
	}

	return nil
}
