package sqlite

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/mattn/go-sqlite3"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
)

func TestUnitNewSqlite(t *testing.T) {

	testCases := []struct {
		name     string
		dbName   string
		dbPath   string
		expected []interface{}
	}{
		{
			name:     "name is empty",
			dbName:   "",
			dbPath:   "ignoreThisForNow",
			expected: []interface{}{&sqlite{}, errInvalidConfiguration},
		},
		{
			name:     "Path is empty",
			dbName:   "ignoreThisForNow",
			dbPath:   "",
			expected: []interface{}{&sqlite{}, errInvalidConfiguration},
		},
		// TODO: Test that exec was called as it should be
	}

	for _, tc := range testCases {
		t.Log(tc.name)

		SQLStorage, err := New(tc.dbPath, tc.dbName, true)

		if !reflect.DeepEqual(tc.expected, []interface{}{SQLStorage, err}) {
			t.Errorf("Expected %v, got %v", tc.expected, SQLStorage)
		}
	}

}

func TestUnitSqliteExecute(t *testing.T) {

	testCases := []struct {
		name      string
		statement string
		arguments []interface{}
		output    error
		withMock  bool
	}{
		{
			name:      "Returns error if DB is undefined",
			statement: "This is a statment",
			arguments: []interface{}{},
			output:    ErrUndefinedDataBase,
		},
		{
			name:      "Returns error if statement is empty",
			statement: "",
			arguments: []interface{}{},
			output:    ErrStatementUndefined,
			withMock:  true,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		s := &sqlite{}
		db, mock, err := sqlmock.New()
		if tc.withMock {
			assert.NoError(t, err)
			s.db = db
		}

		err = s.Execute(tc.statement, tc.arguments...)

		if tc.withMock {
			mock.ExpectationsWereMet()
		}
		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
		db.Close()
	}

	testCasesForBegin := []struct {
		name      string
		statement string
		output    error
	}{
		{
			name:      "Returns error if Begin fails",
			statement: "SELECT * FROM testTable",
			output:    errors.New("Test Error"),
		},
		{
			name:      "Returns no error",
			statement: "SELECT * FROM testTable",
			output:    nil,
		},
	}

	for _, tc := range testCasesForBegin {
		t.Log(tc.name)
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		s := &sqlite{db}

		if tc.output != nil {
			mock.ExpectBegin().WillReturnError(tc.output)
		} else {
			mock.ExpectBegin()
			mock.ExpectExec("^SELECT (.+) FROM testTable").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
		}

		err = s.Execute(tc.statement)

		mock.ExpectationsWereMet()
		if !reflect.DeepEqual(tc.output, err) {
			t.Errorf("Expected '%v', got '%v'", tc.output, err)
		}
		db.Close()
	}
}

func TestUnitsqliteQuery(t *testing.T) {

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
		var dbh *sqlite
		var mock sqlmock.Sqlmock
		var dbConnMock *sql.DB
		var err error
		if tc.hasDatabase {
			dbConnMock, mock, err = sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectQuery("^SELECT (.+) FROM testTable").WillReturnError(sqlite3.ErrError)

			dbh = &sqlite{db: dbConnMock}

		} else {
			dbh = &sqlite{}
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
