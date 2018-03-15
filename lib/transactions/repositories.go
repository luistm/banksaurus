package transactions

import (
	"fmt"
	"strings"

	"github.com/luistm/banksaurus/lib"

	"github.com/shopspring/decimal"

	"github.com/luistm/banksaurus/lib/customerrors"
	"github.com/luistm/banksaurus/lib/sellers"
	"errors"
)

// NewRepository creates a repository for transactions
func NewRepository(storage CSVHandler) *repository {
	return &repository{storage: storage}
}

// CSVHandler to handle csv files
type CSVHandler interface {
	Lines() ([][]string, error)
}

type repository struct {
	storage      CSVHandler
	transactions []lib.Entity
}

func (r *repository) Save(t lib.Entity) error{
	// TODO: Implement this
	return errors.New("Not implemented")
}

func (r *repository) Get(s string)(lib.Entity, error){
	// TODO: Implement this
	return &Transaction{}, errors.New("Not implemented")
}

func (r *repository) GetAll() ([]lib.Entity, error) {

	if r.storage == nil {
		return []lib.Entity{}, customerrors.ErrInfrastructureUndefined
	}

	lines, err := r.storage.Lines()
	if err != nil {
		return []lib.Entity{}, &customerrors.ErrInfrastructure{Msg: err.Error()}
	}

	// TODO: Validate if Lines() output is the expected one
	// r.validateLines(lines)
	// if err != nil{}
	err = r.buildTransactions(lines[5 : len(lines)-2])
	if err != nil {
		return []lib.Entity{}, &customerrors.ErrInfrastructure{Msg: err.Error()}
	}
	// log.Println(len(r.transactions))

	return r.transactions, nil
}

func decimalFromStringWithComma(stringWithComa string) (decimal.Decimal, error) {
	parsedField := strings.Replace(stringWithComa, ".", "", -1)
	parsedField = strings.Replace(parsedField, ",", ".", -1)
	return decimal.NewFromString(parsedField)
}

func (r *repository) buildTransactions(lines [][]string) error {

	for i, line := range lines {

		// TODO: Handle credit
		if line[3] == "" {
			continue
		}

		value, err := decimalFromStringWithComma(line[3])
		if err != nil {
			return fmt.Errorf("failed to create decimal from string: %s", err.Error())
		}

		slug := strings.TrimSuffix(line[2], " ")
		t := &Transaction{
			id:     uint64(i),
			value:  &value,
			Seller: sellers.New(slug, slug),
		}
		r.transactions = append(r.transactions, t)
	}

	return nil
}
