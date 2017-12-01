package transactions

import (
	"fmt"
	"strings"
	"time"

	"github.com/luistm/go-bank-cli/lib/categories"
	"github.com/luistm/go-bank-cli/lib/sellers"
	"github.com/shopspring/decimal"
)

// Fetcher to fetch transactions
type Fetcher interface {
	GetAll() ([]*Transaction, error)
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
	value    string
	s        *sellers.Seller
	c        *categories.Category
	isCredit bool
	date     time.Time
}

func (t *Transaction) String() string {
	return fmt.Sprintf("%s %s", t.Value(), t.s)
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
	parsedField := strings.Replace(t.value, ".", "", -1)
	parsedField = strings.Replace(parsedField, ",", ".", -1)

	retValue, _ := decimal.NewFromString(parsedField)

	return &retValue
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

// Report is a set of transactions
// TODO: A report can have a time range
type Report struct {
	transactions []*Transaction
}

func (r *Report) String() string {
	s := []string{}
	for _, t := range r.transactions {
		s = append(s, t.String())
	}

	return strings.Join(s, "\n")
}
