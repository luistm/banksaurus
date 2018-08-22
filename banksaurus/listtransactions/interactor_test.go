package listtransactions_test

import (
	"errors"
	"github.com/luistm/banksaurus/banksaurus/listtransactions"
	"github.com/luistm/banksaurus/money"
	"github.com/luistm/banksaurus/seller"
	"github.com/luistm/banksaurus/transaction"
	"github.com/luistm/testkit"
	"testing"
	"time"
)

type presenterStub struct {
	seller string
	value  *money.Money
	err    error
}

func (p *presenterStub) Present(data []map[string]*money.Money) error {
	if p.err != nil {
		return p.err
	}

	for key, value := range data[0] {
		p.seller = key
		p.value = value
	}

	return nil
}

func (p *presenterStub) Seller() string {
	return p.seller
}

func (p *presenterStub) Value() string {
	if p.value != nil {
		return p.value.String()
	}

	return ""
}

type transactionRepository struct {
	transactions []*transaction.Entity
	err          error
}

func (tr *transactionRepository) GetAll() ([]*transaction.Entity, error) {
	if tr.err != nil {
		return []*transaction.Entity{}, tr.err
	}

	return tr.transactions, nil
}

func TestUnitNewReport(t *testing.T) {

	_, err := listtransactions.NewInteractor(nil, nil)
	testkit.AssertEqual(t, err, listtransactions.ErrPresenterUndefined)

	_, err = listtransactions.NewInteractor(&presenterStub{}, nil)
	testkit.AssertEqual(t, err, listtransactions.ErrRepositoryUndefined)
}

func TestUnitReport(t *testing.T) {

	s1, err := seller.New("Seller1", "")
	testkit.AssertIsNil(t, err)
	v1, err := money.NewMoney(1234)
	testkit.AssertIsNil(t, err)
	t1, err := transaction.New(1, time.Now(), s1, v1)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name         string
		presenter    *presenterStub
		repository   *transactionRepository
		err          error
		outputSeller string
		outputValue  string
	}{
		{
			name:      "Response has expected data",
			presenter: &presenterStub{},
			repository: &transactionRepository{
				transactions: []*transaction.Entity{t1},
			},
			outputSeller: s1.ID(),
			outputValue:  v1.String(),
		},
		{
			name:      "Returns error if repository returns error",
			presenter: &presenterStub{},
			repository: &transactionRepository{
				transactions: []*transaction.Entity{},
				err:          errors.New("test error"),
			},
			err: errors.New("test error"),
		},
		{
			name:      "Returns error if presenterStub returns error",
			presenter: &presenterStub{err: errors.New("test error")},
			repository: &transactionRepository{
				transactions: []*transaction.Entity{t1},
			},
			err: errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i, err := listtransactions.NewInteractor(tc.presenter, tc.repository)
			testkit.AssertIsNil(t, err)

			err = i.Execute()

			testkit.AssertEqual(t, tc.err, err)
			testkit.AssertEqual(t, tc.outputSeller, tc.presenter.Seller())
			testkit.AssertEqual(t, tc.outputValue, tc.presenter.Value())
		})
	}
}
