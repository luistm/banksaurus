package infrastructure

import (
	"database/sql"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUnitExecutesStatement(t *testing.T) {

	name := "Returns error if sql database is not defined"
	dbh := &DatabaseHandler{}
	err := dbh.Execute("SELECT * FROM testTable")
	assert.EqualError(t, err, ErrDataBaseConnUndefined.Error(), name)

	name = "Returns error if database returns error"
	dbConnMock, mock, err := sqlmock.New()
	assert.NoError(t, err)

	e := &ErrDataBase{"testError"}

	mock.ExpectBegin()
	mock.ExpectExec("^SELECT (.+) FROM testTable").WillReturnError(e)
	dbh = &DatabaseHandler{dbConnMock}

	err = dbh.Execute("SELECT * FROM testTable")

	assert.NoError(t, mock.ExpectationsWereMet(), name)
	assert.EqualError(t, err, e.Error(), name)

	// TODO: Test an insert with values
	// TODO: Test transaction begin error
	// TODO: Test transaction commit error
}

func TestUnitQuery(t *testing.T) {

	testCases := []struct {
		name          string
		errorExpected bool
		query         string
		hasDatabase   bool
	}{
		{
			name:          "Returns error if database not defined",
			errorExpected: true,
			query:         "",
			hasDatabase:   false,
		},
		{
			name:          "Returns error if database query returns error",
			errorExpected: true,
			query:         "SELECT * FROM testTable",
			hasDatabase:   true,
		},
	}

	for _, tc := range testCases {

		// Setup
		var dbh *DatabaseHandler
		var mock sqlmock.Sqlmock
		var dbConnMock *sql.DB
		var err error
		if tc.hasDatabase {
			dbConnMock, mock, err = sqlmock.New()
			assert.NoError(t, err)

			e := &ErrDataBase{"testError"}
			mock.ExpectQuery("^SELECT (.+) FROM testTable").WillReturnError(e)

			dbh = &DatabaseHandler{Database: dbConnMock}

		} else {
			dbh = &DatabaseHandler{}
		}

		// Call the function being tested
		_, err = dbh.Query(tc.query)

		// Assert result
		if tc.hasDatabase {
			assert.NoError(t, mock.ExpectationsWereMet(), tc.name)
		}

		if tc.errorExpected {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}

}
