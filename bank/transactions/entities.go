package transactions

import (
	"time"

	"github.com/luistm/go-bank-cli/lib/sellers"
	"github.com/shopspring/decimal"
)

type iRepository interface {
	GetAll() ([]*Transaction, error)
}

// New creates a record with some parsed data
func New(r record) *Transaction {

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
	value decimal.Decimal
	s     *sellers.Seller
	// c        *categories.Category
	isCredit bool
	date     time.Time
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
	// parsedField := strings.Replace(t.value, ".", "", -1)
	// parsedField = strings.Replace(parsedField, ",", ".", -1)

	// retValue, _ := decimal.NewFromString(parsedField)

	return nil
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

// Record represents a single line of an importend transaction.
// Data will be imported most likely from a text base format.
// Therefore a record will be the representation of each one of those lines.
type record string
