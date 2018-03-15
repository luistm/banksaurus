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
func New() *Transaction {

	t := &Transaction{}

	// for i := 0; i < len(r); i++ {
	// 	t.value = r[i]
	// 	t.Seller = r[2]

	// 	// Expense
	// 	if i == 3 && r[i] != "" {
	// 		t.TransactionType = isDEBT
	// 		return t
	// 	}

	// 	// Credit
	// 	if i == 4 && r[i] != "" {
	// 		t.TransactionType = isCREDIT
	// 		return t
	// 	}

	// 	// Transaction Date
	// 	if i == 0 {
	// 		t.date = r[i]
	// 	}
	// }

	return t
}

// Transaction is a money movement
type Transaction struct {
	id       uint64
	value    *decimal.Decimal
	seller   *sellers.Seller
	isCredit bool
	date     time.Time
}

// ID ...
func (t *Transaction) ID() string {
	return strconv.FormatUint(t.id, 10)
}

func (t *Transaction) String() string {
	return fmt.Sprintf("%s %s", t.Value(), t.seller)
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
