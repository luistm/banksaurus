package loadcsv_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/bankservices"
	"github.com/luistm/banksaurus/bankservices/loadcsv"
	"github.com/luistm/banksaurus/bankservices/transaction"
	"github.com/luistm/testkit"
)

type fileStub struct{}

func (s *fileStub) Lines() ([][]string, error) {
	lines := [][]string{
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{"25-10-2017", "25-10-2017", "COMPRA CAFETARIA HEAR ", "4,30", "", "233,86", "233,86"},
		[]string{},
		[]string{},
	}
	return lines, nil
}

func (s *fileStub) Execute(stmt string, args ...interface{}) error {
	return nil
}

func (s *fileStub) Query(statement string, args ...interface{}) (banklib.Rows, error) {
	return nil, errors.New("this test error should not be happening")
}

type dbStub struct {
	t *testing.T
}

func (dbs *dbStub) Query(statement string, args ...interface{}) (banklib.Rows, error) {
	db, _, err := sqlmock.New()
	testkit.AssertIsNil(dbs.t, err)

	rows, err := db.Query("SELECT SOMETHING")
	testkit.AssertIsNil(dbs.t, err)

	return rows, nil
}

func (dbs *dbStub) Execute(statement string, args ...interface{}) error {
	return errors.New("this should not being called on this tests")
}

func TestIntegrationLoadData(t *testing.T) {

	service := loadcsv.New(&fileStub{}, &fileStub{})
	err := service.Execute()
	testkit.AssertIsNil(t, err)

	presenter := &bankservices.PresenterMock{}
	infrastructure := &dbStub{t}

	service, err = transaction.New(infrastructure, presenter)
	testkit.AssertIsNil(t, err)

	err = service.Execute()
	testkit.AssertIsNil(t, err)

	// TODO: assert presenter has correct data
}
