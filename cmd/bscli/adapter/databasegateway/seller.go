package databasegateway

import (
	"github.com/luistm/banksaurus/seller"
	"github.com/mattn/go-sqlite3"
)

func (r *Repository) GetByID(id string) (*seller.Entity, error) {

	selectStatement := "SELECT * FROM seller WHERE slug = ?"
	rows, err := r.db.Query(selectStatement, id)
	var slug string
	var name string

	for rows.Next() {
		err := rows.Scan(&slug, &name)
		if err != nil {
			return &seller.Entity{}, err
		}

		break // expecting one row max
	}

	s, err := seller.New(slug, "")
	if err != nil {
		return &seller.Entity{}, err
	}

	err = rows.Err()
	if err != nil {
		return &seller.Entity{}, err
	}

	err = rows.Close()
	if err != nil {
		return &seller.Entity{}, err
	}

	return s, nil
}

func (r *Repository) UpdateSeller(seller *seller.Entity) error {

	updateStatement := "UPDATE seller SET name = ? WHERE slug = ?"
	_, err := r.db.Exec(updateStatement, seller.Name(), seller.ID())
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAll() ([]*seller.Entity, error) {

	selectStatement := "SELECT * FROM seller"
	rows, err := r.db.Query(selectStatement)
	if err != nil {
		return []*seller.Entity{}, err
	}

	sellers := []*seller.Entity{}

	for rows.Next() {
		var id string
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			return []*seller.Entity{}, err
		}

		s, err := seller.New(id, name)
		if err != nil {
			return []*seller.Entity{}, err
		}

		sellers = append(sellers, s)
	}

	err = rows.Err()
	if err != nil {
		return []*seller.Entity{}, err
	}

	err = rows.Close()
	if err != nil {
		return []*seller.Entity{}, err
	}

	return sellers, nil
}

// Saves seller to the database
func (r *Repository) Save(seller *seller.Entity) error {

	insertStatement := "INSERT INTO seller(slug, name) VALUES (?, ?)"
	_, err := r.db.Exec(insertStatement, seller.ID(), "")
	if err != nil {
		// Ignore unique
		pqErr, ok := err.(sqlite3.Error)
		if ok && pqErr.Code == sqlite3.ErrConstraint {
			// Should it return the error?
			// Maybe update the name, if needed?
			return nil
		}
		return err
	}

	return nil
}
