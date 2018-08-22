package transactiongateway

import (
	"database/sql"
	"errors"
	"github.com/luistm/banksaurus/seller"
	"github.com/luistm/banksaurus/transaction"
	"github.com/mattn/go-sqlite3"
	"log"
	"strconv"
	"strings"
	"time"
)

// ErrDatabaseUndefined ...
var ErrDatabaseUndefined = errors.New("database is not defined")

// NewTransactionRepository creates a new seller repository instance
func NewTransactionRepository(db *sql.DB) (*Repository, error) {
	if db == nil {
		return &Repository{}, ErrDatabaseUndefined
	}
	return &Repository{db}, nil
}

// Repository for transactions
type Repository struct {
	db *sql.DB
}

// GetAll returns all transactions
func (r *Repository) GetAll() ([]*transaction.Entity, error) {

	statement := `SELECT * FROM "transaction"`
	rows, err := r.db.Query(statement)
	if err != nil {
		return []*transaction.Entity{}, err
	}
	defer rows.Close()

	transactions := []*transaction.Entity{}

	for rows.Next() {
		var id uint64
		var sellerID string
		var value int64

		err := rows.Scan(&id, &sellerID, &value)
		if err != nil {
			return []*transaction.Entity{}, err
		}

		m, err := transaction.NewMoney(value)
		if err != nil {
			return []*transaction.Entity{}, err
		}

		s, err := seller.New(sellerID, "")
		if err != nil {
			return []*transaction.Entity{}, err
		}

		tr, err := transaction.New(id, time.Now(), s, m)
		if err != nil {
			return []*transaction.Entity{}, err
		}

		transactions = append(transactions, tr)
	}

	err = rows.Err()
	if err != nil {
		return []*transaction.Entity{}, err
	}

	return transactions, nil
}

func transactionFromline(line []string) (time.Time, string, *transaction.Money, error) {

	s := strings.TrimSpace(line[2])

	// If not a debt, then is a credit
	isDebt := true
	valueString := line[3]
	if line[4] != "" {
		valueString = line[4]
		isDebt = false
	}

	valueString = strings.Replace(valueString, ",", "", -1)
	valueString = strings.Replace(valueString, ".", "", -1)
	value, err := strconv.ParseInt(valueString, 10, 64)
	if err != nil {
		return time.Time{}, "", &transaction.Money{}, err
	}

	date, err := time.Parse("02-01-2006", line[0])
	if err != nil {
		return time.Time{}, "", &transaction.Money{}, err
	}

	if isDebt {
		value = value * -1
	}

	m, err := transaction.NewMoney(value)
	if err != nil {
		return time.Time{}, "", &transaction.Money{}, err
	}

	return date, s, m, nil
}

// NewFromLine adds a new transaction given it's raw line
func (r *Repository) NewFromLine(line []string) error {

	_, sellerID, m, err := transactionFromline(line)
	if err != nil {
		return err
	}

	s, err := seller.New(sellerID, "")
	if err != nil {
		return err
	}

	err = r.save(s)
	if err != nil {
		return err
	}

	statement := `INSERT INTO "transaction" (seller, amount ) VALUES (?,?)`
	_, err = r.db.Exec(statement, sellerID, m.Value())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Saves seller to the database
func (r *Repository) save(seller *seller.Entity) error {

	insertStatement := "INSERT INTO seller(slug, name) VALUES (?, ?)"
	_, err := r.db.Exec(insertStatement, seller.ID(), "")
	if err != nil {
		// Ignore unique
		pqErr, ok := err.(sqlite3.Error)
		if ok && pqErr.Code == sqlite3.ErrConstraint {
			// Should it return the error?
			// Maybe update the name, if needed?
			return nil
		}
		return err
	}

	return nil
}
