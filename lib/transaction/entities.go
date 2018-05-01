package transaction

import (
	"fmt"
	"strconv"
	"time"

	"github.com/luistm/banksaurus/lib"

	"errors"
	"strings"

	"github.com/luistm/banksaurus/lib/customerrors"
	"github.com/luistm/banksaurus/lib/seller"
	"github.com/shopspring/decimal"
)

// Fetcher to fetch transaction
type Fetcher interface {
	GetAll() ([]lib.Entity, error)
}

// New creates a record with some parsed data
func New(s *seller.Seller, value string) (*Transaction, error) {

	v := value
	if v == "" {
		v = "0"
	}
	valueDecimal, err := decimal.NewFromString(v)
	if err != nil {
		return &Transaction{}, err
	}

	return &Transaction{Seller: s, isCredit: false, value: &valueDecimal}, nil
}

// Transaction is a money movement
type Transaction struct {
	id       uint64
	value    *decimal.Decimal
	Seller   *seller.Seller
	isCredit bool
	date     time.Time
}

// ID of a transaction
func (t *Transaction) ID() string {
	return strconv.FormatUint(t.id, 10)
}

// String to satisfy the fmt.Stringer interface
func (t *Transaction) String() string {
	return fmt.Sprintf("%s %s", t.Value(), t.Seller)
}

// IsDebt returns true if a transaction is a debt
func (t *Transaction) IsDebt() bool {
	if t.isCredit {
		return false
	}
	return true
}

// Value returns the field value parsed for money
func (t *Transaction) Value() *decimal.Decimal {
	return t.value
}

// NewRepository creates a Transactions for transaction
func NewRepository(storage CSVHandler) *Transactions {
	return &Transactions{storage: storage}
}

// CSVHandler to handle csv files
type CSVHandler interface {
	Lines() ([][]string, error)
}

// Transactions repository
type Transactions struct {
	storage      CSVHandler
	transactions []lib.Entity
}

// Save to save a transaction
func (r *Transactions) Save(t lib.Entity) error {
	// TODO: Implement this
	return errors.New("save not implemented")
}

// Get to fetch a single transaction
func (r *Transactions) Get(s string) (lib.Entity, error) {
	// TODO: Implement this
	return &Transaction{}, errors.New("get not implemented")
}

// GetAll to fetch all transactions
func (r *Transactions) GetAll() ([]lib.Entity, error) {
	// TODO: Should return an iterator

	if r.storage == nil {
		return []lib.Entity{}, customerrors.ErrInfrastructureUndefined
	}

	lines, err := r.storage.Lines()
	if err != nil {
		return []lib.Entity{}, &customerrors.ErrInfrastructure{Msg: err.Error()}
	}

	// TODO: Validate if Lines() output is the expected one

	err = r.buildTransactions(lines[5 : len(lines)-2])
	if err != nil {
		return []lib.Entity{}, &customerrors.ErrInfrastructure{Msg: err.Error()}
	}

	return r.transactions, nil
}

func decimalFromStringWithComma(stringWithComa string) (decimal.Decimal, error) {
	parsedField := strings.Replace(stringWithComa, ".", "", -1)
	parsedField = strings.Replace(parsedField, ",", ".", -1)
	return decimal.NewFromString(parsedField)
}

func (r *Transactions) buildTransactions(lines [][]string) error {

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
			Seller: seller.New(slug, slug),
		}
		r.transactions = append(r.transactions, t)
	}

	return nil
}
