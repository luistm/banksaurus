package sqlite_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/luistm/banksaurus/next/adapter/sqlite"
	"github.com/luistm/testkit"
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

//func TestUnitSellerRepositorySave(t *testing.T){
//
//	testCases := []struct{
//		name string
//		input *seller.Entity
//		expectedErr error
//	}{
//		{
//			name: "Saves seller to database",
//		},
//	}
//
//	for _, tc := range testCases{
//		t.Run(tc.name, func(t *testing.T) {
//			db, mock, err := sqlmock.New()
//			if err != nil {
//				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//			}
//			defer db.Close()
//
//			//mock....
//			//Ler database applications Go
//
//			r, err := sqlite.NewSellerRepository(db)
//			testkit.AssertIsNil(t, err)
//
//			err = r.Save(tc.input)
//
//			testkit.AssertIsNil(t, mock.ExpectationsWereMet())
//			testkit.AssertEqual(t, tc.expectedErr, err)
//		})
//	}
//}
//
