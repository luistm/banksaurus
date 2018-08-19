package databasegateway_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/luistm/banksaurus/cmd/bscli/adapter/databasegateway"
	"github.com/luistm/testkit"
	"testing"
)

func TestUnitSellerRepositoryNew(t *testing.T) {

	t.Run("Returns error if database is nil", func(t *testing.T) {
		_, err := databasegateway.NewSellerRepository(nil)
		testkit.AssertEqual(t, databasegateway.ErrDatabaseUndefined, err)
	})

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("Does not return error if receives a database", func(t *testing.T) {
		_, err := databasegateway.NewSellerRepository(db)
		testkit.AssertIsNil(t, err)
	})
}
