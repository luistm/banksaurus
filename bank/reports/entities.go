package reports

import (
	"strings"

	"github.com/luistm/go-bank-cli/bank/transactions"
)

// Report ...
type Report struct {
	transactions []*transactions.Transaction
}

func (r *Report) String() string {
	s := []string{}
	for _, t := range r.transactions {
		s = append(s, t.String())
	}

	return strings.Join(s, "\n")
}
