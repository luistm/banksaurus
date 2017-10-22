package categories

import (
	"errors"
	"fmt"

	"github.com/luistm/go-bank-cli/entities"
)

var errInvalidCategory = errors.New("Invalid category")

const insertStatement string = "INSERT INTO categories(name) VALUES(?)"

// repository allows us the save a read categories from a repository
type repository struct {
	SQLStorage entities.SQLDatabaseHandler
}

// Save to persist a category
func (r *repository) Save(c *Category) error {

	if c == nil || c.Name == "" {
		return errInvalidCategory
	}

	if r.SQLStorage == nil {
		return entities.ErrInfrastructureUndefined
	}

	if err := r.SQLStorage.Execute(insertStatement, c.Name); err != nil {
		return &entities.ErrInfrastructure{Msg: err.Error()}
	}

	return nil
}

// Get fetches a category by name
func (r *repository) Get(name string) (*Category, error) {
	statement := "SELECT * FROM categories WHERE name=?"
	_, err := r.SQLStorage.Query(statement)
	if err != nil {
		return &Category{}, fmt.Errorf("Database failure: %s", err)
	}

	return &Category{}, nil
}

// GetAll fetches all categories
func (r *repository) GetAll() ([]*Category, error) {
	statement := "SELECT * FROM categories"
	rows, err := r.SQLStorage.Query(statement)
	if err != nil {
		return []*Category{}, fmt.Errorf("Database failure: %s", err)
	}

	categories := []*Category{}

	for rows.Next() {
		var id int
		var cat string
		err := rows.Scan(&id, &cat)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &Category{Name: cat})
	}

	return categories, nil
}
