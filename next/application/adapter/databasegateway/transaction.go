package databasegateway

import (
	"github.com/luistm/banksaurus/next/entity/seller"
	"github.com/luistm/banksaurus/next/entity/transaction"
	"strconv"
	"strings"
	"time"
)

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

	return nil
	_, sellerID, _, err := transactionFromline(line)
	if err != nil {
		return err
	}

	s, err := seller.New(sellerID, "")
	if err != nil {
		return err
	}

	err = r.Save(s)
	if err != nil {
		return err
	}

	// Save to database
	// TODO: Add function to load transactions into the database
	//       transaction, err := TransactionFactory(record)
	//       Transactions should now have an ID, a sellerID, a value and a date

	// Return the transactions after adding the ir coming from the database
	//CREATE TABLE IF NOT EXISTS transactions
	//(
	//	ID int NOT NULL PRIMARY KEY,
	//	SELLER_ID int NOT NULL,
	//	AMOUNT int DEFAULT 0,
	//	TYPE
	//BALANCE int NOT NULL
	//);

	return nil
}
