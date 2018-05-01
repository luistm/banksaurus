package bank_test

import (
	"testing"
	"github.com/luistm/banksaurus/bank"
	"github.com/luistm/banksaurus/elib/testkit"
)

func TestIntegrationTransactionsShow(t *testing.T) {

	transactionShowUseCase, err := bank.New()
	testkit.AssertIsNil(t, err)

	err = transactionShowUseCase.Execute()
	testkit.AssertIsNil(t, err)


}
