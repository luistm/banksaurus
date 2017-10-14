package reports

import (
	"github.com/luistm/go-bank-cli/lib/categories"
	"github.com/luistm/go-bank-cli/lib/descriptions"
)

// Transaction is a single movement in an account
type Transaction struct {
	c *categories.Category
	d *descriptions.Description
	v int64
}

// Report ...
type Report struct {
	transactions []*Transaction
}
