package report_test

import (
	"errors"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"github.com/luistm/banksaurus/next/interactor/report"
	"github.com/luistm/testkit"
	"strconv"
	"testing"
)

type presenter struct {
	rawData map[string]string
	err     error
}

func (p *presenter) Data() map[string]string {
	return p.rawData
}

func (p *presenter) Present(ts []*transaction.Entity) error {
	if p.err != nil {
		return p.err
	}

	p.rawData = map[string]string{"id": strconv.FormatUint(ts[0].ID(), 10)}
	return nil
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

func TestNewReport(t *testing.T) {

	_, err := report.NewInteractor(nil, nil)
	testkit.AssertEqual(t, err, report.ErrPresenterUndefined)

	_, err = report.NewInteractor(&presenter{}, nil)
	testkit.AssertEqual(t, err, report.ErrRepositoryUndefined)
}

func TestReport(t *testing.T) {

	rq1, err := report.NewRequest()
	testkit.AssertIsNil(t, err)

	t1, err := transaction.New()
	testkit.AssertIsNil(t, err)

	testCases := []struct {
		name       string
		input      *report.Request
		presenter  *presenter
		repository *transactionRepository
		err        error
		output     map[string]string
	}{
		{
			name:      "Response has expected data",
			input:     rq1,
			presenter: &presenter{},
			repository: &transactionRepository{
				transactions: []*transaction.Entity{t1},
			},
			output: map[string]string{"id": strconv.FormatUint(t1.ID(), 10)},
		},
		{
			name:      "Returns error if repository returns error",
			input:     rq1,
			presenter: &presenter{},
			repository: &transactionRepository{
				transactions: []*transaction.Entity{},
				err:          errors.New("test error"),
			},
			err: report.ErrRepository.AppendError(errors.New("test error")),
		},
		{
			name:      "Returns error if presenter returns error",
			input:     rq1,
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

			err = i.Execute(tc.input)

			testkit.AssertEqual(t, tc.err, err)
			testkit.AssertEqual(t, tc.output, tc.presenter.Data())
		})
	}
}
