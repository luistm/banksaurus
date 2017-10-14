package sqlite

import (
	"database/sql"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
)

func TestUnitNew(t *testing.T) {

	testCases := []struct {
		name          string
		dbName        string
		dbPath        string
		errorExpected bool
	}{
		{
			name:          "Name is empty",
			dbName:        "",
			dbPath:        "ignoreThisForNow",
			errorExpected: true,
		},
		{
			name:          "Path is empty",
			dbName:        "ignoreThisForNow",
			dbPath:        "",
			errorExpected: true,
		},
		{
			name:          "Path does not exist",
			dbName:        "ignoreThisForNow",
			dbPath:        "./ThisPathDoesNotExist",
			errorExpected: false,
		},
	}

	for _, tc := range testCases {
		// TODO: test that the correct interface is returned
		_, err := New(tc.dbPath, tc.dbName, true)

		if tc.errorExpected {
			assert.Error(t, err, tc.name)
		} else {
			assert.NoError(t, err, tc.name)
		}
	}

}

func TestUnitSqliteExecute(t *testing.T) {

	name := "Returns error if sql database is not defined"
	dbh := &Sqlite{}
	err := dbh.Execute("SELECT * FROM testTable")
	assert.EqualError(t, err, errConnectionIsNil.Error(), name)

	name = "Returns error if database returns error"
	dbConnMock, mock, err := sqlmock.New()
	assert.NoError(t, err)

	e := &ErrDataBase{"testError"}

	mock.ExpectBegin()
	mock.ExpectExec("^SELECT (.+) FROM testTable").WillReturnError(e)
	dbh = &Sqlite{dbConnMock}

	err = dbh.Execute("SELECT * FROM testTable")

	assert.NoError(t, mock.ExpectationsWereMet(), name)
	assert.EqualError(t, err, e.Error(), name)

	// TODO: Test an insert with values
	// TODO: Test transaction begin error
	// TODO: Test transaction commit error
}

func TestUnitSqliteQuery(t *testing.T) {

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
		var dbh *Sqlite
		var mock sqlmock.Sqlmock
		var dbConnMock *sql.DB
		var err error
		if tc.hasDatabase {
			dbConnMock, mock, err = sqlmock.New()
			assert.NoError(t, err)

			e := &ErrDataBase{"testError"}
			mock.ExpectQuery("^SELECT (.+) FROM testTable").WillReturnError(e)

			dbh = &Sqlite{db: dbConnMock}

		} else {
			dbh = &Sqlite{}
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
