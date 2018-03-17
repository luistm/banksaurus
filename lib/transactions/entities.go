package transactions

import (
	"fmt"
	"strconv"
	"time"

	"github.com/luistm/banksaurus/lib"

	"github.com/luistm/banksaurus/lib/sellers"
	"github.com/shopspring/decimal"
)

// Fetcher to fetch transactions
type Fetcher interface {
	GetAll() ([]lib.Entity, error)
}

// New creates a record with some parsed data
func New(s *sellers.Seller, value string) (*Transaction, error) {

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
	Seller   *sellers.Seller
	isCredit bool
	date     time.Time
}

// ID ...
func (t *Transaction) ID() string {
	return strconv.FormatUint(t.id, 10)
}

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

// IsFromThisMonth checks if a transaction is from the current month
func (t *Transaction) IsFromThisMonth() bool {

	// This code sucks so much.... i'm sure there is a better way, i'm just to lazy now...
	// splittedDate := strings.Split(t.date, "-")
	// year, _ := strconv.Atoi(splittedDate[2])
	// month, _ := strconv.Atoi(splittedDate[1])

	// if time.Month(month) == time.Now().Month() &&
	// 	year == time.Now().Year() {
	// 	return true
	// }
	return false
}
