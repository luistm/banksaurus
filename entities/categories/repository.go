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
func (r *repository) Save(ent entities.Entity) error {

	c := ent.(*Category)
	if c == nil || c.name == "" {
		return errInvalidCategory
	}

	if r.SQLStorage == nil {
		return entities.ErrInfrastructureUndefined
	}

	if err := r.SQLStorage.Execute(insertStatement, c.name); err != nil {
		return &entities.ErrInfrastructure{Msg: err.Error()}
	}

	return nil
}

// Get fetches a category by name
func (r *repository) Get(name string) (entities.Entity, error) {
	statement := "SELECT * FROM categories WHERE name=?"
	_, err := r.SQLStorage.Query(statement)
	if err != nil {
		return &Category{}, fmt.Errorf("Database failure: %s", err)
	}

	return &Category{}, nil
}

// GetAll fetches all categories
func (r *repository) GetAll() ([]entities.Entity, error) {
	statement := "SELECT * FROM categories"
	rows, err := r.SQLStorage.Query(statement)
	if err != nil {
		return []entities.Entity{}, fmt.Errorf("Database failure: %s", err)
	}

	categories := []entities.Entity{}

	for rows.Next() {
		var id int
		var cat string
		err := rows.Scan(&id, &cat)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &Category{name: cat})
	}

	return categories, nil
}
