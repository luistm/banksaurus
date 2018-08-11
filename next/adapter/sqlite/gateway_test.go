package sqlite_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/luistm/banksaurus/next/adapter/sqlite"
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/testkit"
	"github.com/mattn/go-sqlite3"
	"testing"
)

func TestUnitSellerRepositoryNew(t *testing.T) {

	t.Run("Returns error if database is nil", func(t *testing.T) {
		_, err := sqlite.NewSellerRepository(nil)
		testkit.AssertEqual(t, sqlite.ErrDatabaseUndefined, err)
	})

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("Does not return error if receives a database", func(t *testing.T) {
		_, err := sqlite.NewSellerRepository(db)
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitSellerGetAll(t *testing.T) {

	testCases := []struct {
		name            string
		expectedSellers []*seller.Entity
		expectedError   error
	}{
		{},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}

func TestUnitSellerRepositorySave(t *testing.T) {

	s1, err := seller.NewFromID("SellerID")
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name         string
		input        *seller.Entity
		expectedErr  error
		sqlMockError error
	}{
		{
			name:  "Saves seller to database",
			input: s1,
		},
		{
			name:         "Returns error if db returns error",
			input:        s1,
			expectedErr:  errors.New("this is a database error"),
			sqlMockError: errors.New("this is a database error"),
		},
		{
			name:         "Returns success if error is: UNIQUE constraint failed",
			input:        s1,
			expectedErr:  nil,
			sqlMockError: sqlite3.Error{Code: sqlite3.ErrConstraint},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectExec("INSERT INTO seller").
				WithArgs(tc.input.ID()).
				WillReturnError(tc.sqlMockError).
				WillReturnResult(sqlmock.NewResult(1, 1))

			r, err := sqlite.NewSellerRepository(db)
			testkit.AssertIsNil(t, err)

			err = r.Save(tc.input)

			testkit.AssertIsNil(t, mock.ExpectationsWereMet())
			testkit.AssertEqual(t, tc.expectedErr, err)
		})
	}
}
