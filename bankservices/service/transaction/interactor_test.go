package transaction_test

import (
	"testing"

	"github.com/luistm/banksaurus/bank/usecase/transaction"
	"github.com/luistm/banksaurus/elib/testkit"
)

//type infrastructureStub struct{
//	t *testing.T
//}
//
//func (is *infrastructureStub) Execute(q *lib.Inquirer) (lib.Rows, error){
//
//	db, _, err := sqlmock.New()
//	testkit.AssertIsNil(is.t, err)
//
//	rows, err := db.Query("SELECT SOMETHING")
//	testkit.AssertIsNil(is.t, err)
//
//	return rows, nil
//
//}

func TestIntegrationTransactionsShow(t *testing.T) {

	//is := infrastructureStub{t}

	transactionShowUseCase, err := transaction.New()
	testkit.AssertIsNil(t, err)

	err = transactionShowUseCase.Execute()
	testkit.AssertIsNil(t, err)

	t.Error("Work in progress")

}
