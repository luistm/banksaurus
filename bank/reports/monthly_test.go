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
		mockRepository *mockRepository
		mockReturn     []interface{}
		expectedReport *Report
		expectedError  error
	}{
		{
			name:           "Returns error if respository is not defined",
			mockRepository: nil,
			mockReturn:     []interface{}{},
			expectedReport: &Report{},
			expectedError:  errors.New("Failed"),
		},
		{
			name:           "Returns error if transactions repository returns error",
			mockRepository: new(mockRepository),
			mockReturn: []interface{}{[]*Transaction{},
				errors.New("repository error")},
			expectedReport: &Report{},
			expectedError:  errors.New("Failed"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i := &Interactor{}
			if tc.mockRepository != nil {
				tc.mockRepository.On("AllTransactions").Return(tc.mockReturn...)
				i.repository = tc.mockRepository
			}

			r, err := i.MonthlyReport()

			if tc.mockRepository != nil {
				tc.mockRepository.AssertExpectations(t)
			}
			if tc.expectedError != nil && err == nil {
				t.Errorf("Expecting %v, got %v", tc.expectedError, err)
			}

			if !reflect.DeepEqual(r, tc.expectedReport) {
				t.Errorf("Expecting %v, got %v", tc.expectedReport, r)
			}

		})
	}

}
