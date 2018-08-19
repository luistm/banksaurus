package sellergateway_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/luistm/banksaurus/cmd/bscli/adapter/sellergateway"
	"github.com/luistm/testkit"
	"testing"
	"github.com/luistm/banksaurus/seller"
	"github.com/mattn/go-sqlite3"
)

func TestUnitSellerRepositoryNew(t *testing.T) {

	t.Run("Returns error if database is nil", func(t *testing.T) {
		_, err := sellergateway.NewSellerRepository(nil)
		testkit.AssertEqual(t, sellergateway.ErrDatabaseUndefined, err)
	})

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("Does not return error if receives a database", func(t *testing.T) {
		_, err := sellergateway.NewSellerRepository(db)
		testkit.AssertIsNil(t, err)
	})
}


func TestUnitSellerGetAll(t *testing.T) {

	s1, err := seller.New("SellerID", "")
	testkit.AssertIsNil(t, err)
	s2, err := seller.New("SellerID", "")
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name            string
		expectedSellers []*seller.Entity
		expectedError   error
		dbError         error
		dbRows          *sqlmock.Rows
	}{
		{
			name:            "Returns sellers",
			expectedSellers: []*seller.Entity{s1, s2},
			dbRows:          sqlmock.NewRows([]string{"seller", "name"}).AddRow(s1.ID(), "").AddRow(s2.ID(), ""),
		},
		{
			name:            "Handles query error",
			expectedSellers: []*seller.Entity{},
			dbError:         errors.New("test error"),
			expectedError:   errors.New("test error"),
		},
		{
			name:            "Handles scan error",
			expectedSellers: []*seller.Entity{},
			dbRows:          sqlmock.NewRows([]string{"seller", "name"}).AddRow(s1.ID(), "").RowError(0, errors.New("test error")),
			expectedError:   errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			r, err := sellergateway.NewSellerRepository(db)
			testkit.AssertIsNil(t, err)

			mock.ExpectQuery("SELECT (.*) FROM seller").WillReturnRows(tc.dbRows).WillReturnError(tc.dbError)

			sellers, err := r.GetAll()

			testkit.AssertIsNil(t, mock.ExpectationsWereMet())
			testkit.AssertEqual(t, tc.expectedError, err)
			testkit.AssertEqual(t, tc.expectedSellers, sellers)

		})
	}
}

func TestUnitSellerRepositorySave(t *testing.T) {

	s1, err := seller.New("SellerID", "")
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
				WithArgs(tc.input.ID(), "").
				WillReturnError(tc.sqlMockError).
				WillReturnResult(sqlmock.NewResult(1, 1))

			r, err := sellergateway.NewSellerRepository(db)
			testkit.AssertIsNil(t, err)

			err = r.Save(tc.input)

			testkit.AssertIsNil(t, mock.ExpectationsWereMet())
			testkit.AssertEqual(t, tc.expectedErr, err)
		})
	}
}
