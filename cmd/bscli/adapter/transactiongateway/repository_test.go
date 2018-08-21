package transactiongateway_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/luistm/banksaurus/cmd/bscli/adapter/transactiongateway"
	"github.com/luistm/testkit"
	"testing"
)

// line := "25-10-2017;25-10-2017;COMPRA CONTINENTE MAI ;77,52;;61,25;61,25;"
// line := ["25-10-2017", "25-10-2017", "COMPRA CONTINENTE MAI ", "77,52", "61,25", "61,25"]

//func TestUnitTransactionFactoryNewTransaction(t *testing.T){
//
//	testCases := []struct{
//
//	}{
//		{
//			a line here
//		}
//	}
//}

func TestUnitRepositoryNew(t *testing.T) {

	db, _, err := sqlmock.New()
	testkit.AssertIsNil(t, err)

	t.Run("Returns error if infrastructure not defined", func(t *testing.T) {
		_, err := transactiongateway.NewTransactionRepository(nil)
		testkit.AssertEqual(t, transactiongateway.ErrDatabaseUndefined, err)
	})

	t.Run("Returns no error", func(t *testing.T) {
		_, err := transactiongateway.NewTransactionRepository(db)
		testkit.AssertIsNil(t, err)
	})
}

func TestUnitRepositoryNewTransaction(t *testing.T) {

}
