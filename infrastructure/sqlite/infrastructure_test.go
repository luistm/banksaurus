package sqlite

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

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
			expected: []interface{}{&Infrastructure{}, errInvalidConfiguration},
		},
		{
			name:     "Path is empty",
			dbName:   "ignoreThisForNow",
			dbPath:   "",
			expected: []interface{}{&Infrastructure{}, errInvalidConfiguration},
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
			statement: "This is a statement",
			arguments: []interface{}{},
			output:    errUndefinedDataBase,
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
		s := &Infrastructure{}
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
			output:    errors.New("test Error"),
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
		s := &Infrastructure{db}

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

// TODO: Complete this test
// func TestUnitSqliteQuery(t *testing.T) {

// 	testCases := []struct {
// 		name     string
// 		input    []interface{}
// 		output   []interface{}
// 		withMock bool
// 	}{
// 		{
// 			name:     "Returns error if database not defined",
// 			input:    []interface{}{"", []interface{}{}},
// 			output:   []interface{}{nil, errUndefinedDataBase},
// 			withMock: false,
// 		},
// 		{
// 			name:     "Returns error if database query returns error",
// 			input:    []interface{}{"SELECT * FROM testTable", []interface{}{}},
// 			output:   []interface{}{nil, sqlite3.ErrError},
// 			withMock: true,
// 		},
// 		// {
// 		// 	name:     "Returns rows from query",
// 		// 	input:    []interface{}{"SELECT * FROM testTable WHERE id=?", []interface{}{1}},
// 		// 	withMock: true,
// 		// },
// 	}

// 	for _, tc := range testCases {
// 		t.Log(tc.name)
// 		dbh := &Infrastructure{}
// 		var mock sqlmock.Sqlmock
// 		var err error
// 		var dbMock *sql.DB
// 		if tc.withMock {
// 			dbMock, mock, err = sqlmock.New()
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			mock.ExpectQuery("^SELECT (.+) FROM testTable").WillReturnError(sqlite3.ErrError)
// 			dbh.db = dbMock
// 		}

// 		rows, err := dbh.Query(tc.input[0].(string), tc.input[1].([]interface{})...)

// 		got := []interface{}{rows, err}
// 		if tc.withMock {
// 			err = mock.ExpectationsWereMet()
// 			if err != nil {
// 				t.Error(err)
// 			}
// 		}
// 		testkit.AssertEqual(t, tc.output, got)
// 	}
// }
