package reports

import "github.com/luistm/go-bank-cli/bank/transactions"

// Report ...
type Report struct {
	transactions []*transactions.Transaction
}
