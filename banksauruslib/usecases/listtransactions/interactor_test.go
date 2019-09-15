package listtransactions_test

import (
	"errors"
	"github.com/luistm/banksaurus/banksauruslib/entities/seller"
	"github.com/luistm/banksaurus/banksauruslib/entities/transaction"
	"github.com/luistm/banksaurus/banksauruslib/usecases/listtransactions"
	"github.com/luistm/testkit"
	"testing"
	"time"
)

type presenter struct {
	seller string
	value  *transaction.Money
	err    error
}

func (p *presenter) Present(data []map[string]*transaction.Money) error {
	if p.err != nil {
		return p.err
	}

	for key, value := range data[0] {
		p.seller = key
		p.value = value
	}

	return nil
}

func (p *presenter) Seller() string {
	return p.seller
}

func (p *presenter) Value() string {
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

	_, err = listtransactions.NewInteractor(&presenter{}, nil)
	testkit.AssertEqual(t, err, listtransactions.ErrRepositoryUndefined)
}

func TestUnitReport(t *testing.T) {

	s1, err := seller.New("Seller1", "")
	testkit.AssertIsNil(t, err)
	v1, err := transaction.NewMoney(1234)
	testkit.AssertIsNil(t, err)
	t1, err := transaction.New(1, time.Now(), s1, v1)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name         string
		presenter    *presenter
		repository   *transactionRepository
		err          error
		outputSeller string
		outputValue  string
	}{
		{
			name:      "Response has expected data",
			presenter: &presenter{},
			repository: &transactionRepository{
				transactions: []*transaction.Entity{t1},
			},
			outputSeller: s1.ID(),
			outputValue:  v1.String(),
		},
		{
			name:      "Returns error if repository returns error",
			presenter: &presenter{},
			repository: &transactionRepository{
				transactions: []*transaction.Entity{},
				err:          errors.New("test error"),
			},
			err: errors.New("test error"),
		},
		{
			name:      "Returns error if presenter returns error",
			presenter: &presenter{err: errors.New("test error")},
			repository: &transactionRepository{
				transactions: []*transaction.Entity{t1},
			},
			err: listtransactions.ErrPresenter.AppendError(errors.New("test error")),
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
