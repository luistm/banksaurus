package categories

import (
	"errors"
	"fmt"
)

var errInvalidCategory = errors.New("Invalid category")

type errInfrastructure struct {
	arg string
}

func (e *errInfrastructure) Error() string {
	return fmt.Sprintf("Infrastructure error: %s", e.arg)
}

const insertStatement string = "INSERT INTO categories(name) VALUES(?)"

// IRow ...
type IRow interface {
	Scan(dest ...interface{})
	Next() bool
}

// IDBHandler ...
type IDBHandler interface {
	Execute(statement string) error
	Query(statement string) (IRow, error)
}

// CategoryRepository allows us the save a read categories from a repository
type CategoryRepository struct {
	dbHandler IDBHandler
}

// Save to persist a category
func (cr *CategoryRepository) Save(c *Category) error {

	if c == nil || c.name == "" {
		return errInvalidCategory
	}

	if err := cr.dbHandler.Execute(insertStatement); err != nil {
		return &errInfrastructure{arg: err.Error()}
	}

	return nil
}

// Get fetches a category by name
func (cr *CategoryRepository) Get(name string) (*Category, error) {
	statement := "SELECT * FROM categories WHERE name=?"
	_, err := cr.dbHandler.Query(statement)
	if err != nil {
		return &Category{}, fmt.Errorf("Failed to get category, database faillure: %s", err)
	}

	// Build category from rows

	return &Category{}, nil
}
