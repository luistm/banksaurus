package categories

import (
	"database/sql"
	"errors"
)

// CategoryRepository allows us the save a read categories from a repository
type CategoryRepository struct {
	db *sql.DB
}

// Save to persist a category
func (cr *CategoryRepository) Save(c *Category) error {

	// Query goes here...
	// INSERT INTO CATEGORY VALUE (name);

	return errors.New("Failed to save category")
}
