package categories

import "errors"

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
	dbHandler *IDBHandler
}

// Save to persist a category
func (cr *CategoryRepository) Save(c *Category) error {

	// Query goes here...
	// INSERT INTO CATEGORY VALUE (name);

	return errors.New("Failed to save category")
}
