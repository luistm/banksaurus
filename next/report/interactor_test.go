package report_test

import (
	"errors"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"github.com/luistm/banksaurus/next/report"
	"github.com/luistm/testkit"
	"testing"
	"time"
)

type presenter struct {
	seller string
	value  int64
	err    error
}

func (p *presenter) Present(data []map[string]int64) error {
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

func (p *presenter) Value() int64 {
	return p.value
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

	_, err := report.NewInteractor(nil, nil)
	testkit.AssertEqual(t, err, report.ErrPresenterUndefined)

	_, err = report.NewInteractor(&presenter{}, nil)
	testkit.AssertEqual(t, err, report.ErrRepositoryUndefined)
}

func TestUnitReport(t *testing.T) {

	s1 := "Seller1"
	v1 := int64(1234)
	t1, err := transaction.New(time.Now(), s1, v1)
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name         string
		presenter    *presenter
		repository   *transactionRepository
		err          error
		outputSeller string
		outputValue  int64
	}{
		{
			name:      "Response has expected data",
			presenter: &presenter{},
			repository: &transactionRepository{
				transactions: []*transaction.Entity{t1},
			},
			outputSeller: s1,
			outputValue:  v1,
		},
		{
			name:      "Returns error if repository returns error",
			presenter: &presenter{},
			repository: &transactionRepository{
				transactions: []*transaction.Entity{},
				err:          errors.New("test error"),
			},
			err: report.ErrRepository.AppendError(errors.New("test error")),
		},
		{
			name:      "Returns error if presenter returns error",
			presenter: &presenter{err: errors.New("test error")},
			repository: &transactionRepository{
				transactions: []*transaction.Entity{t1},
			},
			err: report.ErrPresenter.AppendError(errors.New("test error")),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i, err := report.NewInteractor(tc.presenter, tc.repository)
			testkit.AssertIsNil(t, err)

			err = i.Execute()

			testkit.AssertEqual(t, tc.err, err)
			testkit.AssertEqual(t, tc.outputSeller, tc.presenter.Seller())
			testkit.AssertEqual(t, tc.outputValue, tc.presenter.Value())
		})
	}
}
