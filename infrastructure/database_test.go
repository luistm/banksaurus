package infrastructure

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestExecutesStatement(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

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

	// TODO: Test transaction begin error
	// TODO: Test transaction commit error
}
