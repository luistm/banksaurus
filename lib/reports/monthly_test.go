package reports

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) AllTransactions() ([]*Transaction, error) {
	args := m.Called()
	return args.Get(0).([]*Transaction), args.Error(1)
}

func TestUnitInteractorReportMonthly(t *testing.T) {

	testCases := []struct {
		name           string
		i              *Interactor
		mockRepository *mockRepository
		mockReturn     []interface{}
		expectedReport *Report
		expectedError  error
	}{
		{
			name:           "Returns error if transactions repository returns error",
			i:              &Interactor{},
			mockRepository: new(mockRepository),
			mockReturn:     []interface{}{[]*Transaction{}, errors.New("Repository error")},
			expectedReport: &Report{},
			expectedError:  errors.New("Failed"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			tc.mockRepository.On("AllTransactions").Return(tc.mockReturn...)
			tc.i.Repository = tc.mockRepository

			r, err := tc.i.MonthlyReport()

			tc.mockRepository.AssertExpectations(t)
			if tc.expectedError != nil && err == nil {
				t.Errorf("Expecting %v, got %v", tc.expectedError, err)
			}

			if !reflect.DeepEqual(r, tc.expectedReport) {
				t.Errorf("Expecting %v, got %v", tc.expectedReport, r)
			}

		})
	}

}
