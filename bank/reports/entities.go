package reports

import (
	"github.com/luistm/go-bank-cli/entities/categories"
	"github.com/luistm/go-bank-cli/entities/sellers"
)

// Transaction is a single movement in an account
type Transaction struct {
	c *categories.Category
	s *sellers.Seller
	v int64
}

// Report ...
type Report struct {
	transactions []*Transaction
}
