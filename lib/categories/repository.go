package categories

import (
	"errors"
	"fmt"

	"github.com/luistm/go-bank-cli/lib"
)

var errInvalidCategory = errors.New("Invalid category")

const insertStatement string = "INSERT INTO categories(name) VALUES(?)"

// repository allows us the save a read categories from a repository
type repository struct {
	SQLStorage lib.SQLDatabaseHandler
}

// Save to persist a category
func (r *repository) Save(ent lib.Entity) error {

	c := ent.(*Category)
	if c == nil || c.name == "" {
		return errInvalidCategory
	}

	if r.SQLStorage == nil {
		return lib.ErrInfrastructureUndefined
	}

	if err := r.SQLStorage.Execute(insertStatement, c.name); err != nil {
		return &lib.ErrInfrastructure{Msg: err.Error()}
	}

	return nil
}

// Get fetches a category by name
func (r *repository) Get(name string) (lib.Entity, error) {
	statement := "SELECT * FROM categories WHERE name=?"
	_, err := r.SQLStorage.Query(statement)
	if err != nil {
		return &Category{}, fmt.Errorf("Database failure: %s", err)
	}

	return &Category{}, nil
}

// GetAll fetches all categories
func (r *repository) GetAll() ([]lib.Entity, error) {
	statement := "SELECT * FROM categories"
	rows, err := r.SQLStorage.Query(statement)
	if err != nil {
		return []lib.Entity{}, fmt.Errorf("Database failure: %s", err)
	}

	categories := []lib.Entity{}

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
