package loaddata_test

import (
	"errors"
	"testing"

	"github.com/luistm/banksaurus/banklib"
	"github.com/luistm/banksaurus/bankservices/loaddata"
	"github.com/luistm/testkit"
)

type stub struct{}

func (s *stub) Lines() ([][]string, error) {
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

func (s *stub) Execute(stmt string, args ...interface{}) error {
	return nil
}

func (s *stub) Query(statement string, args ...interface{}) (banklib.Rows, error) {
	return nil, errors.New("this test error should not be happening")
}

func TestIntegrationLoadData(t *testing.T) {
	service := loaddata.New(&stub{}, &stub{})
	err := service.Execute()
	testkit.AssertIsNil(t, err)
}
