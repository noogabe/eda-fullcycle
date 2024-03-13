package database

import (
	"database/sql"
	"fmt"

	"github.com.br/noogabe/eda-fullcycle/consumer/internal/domain/balance"
)

type BalanceDb struct {
	DB *sql.DB
}

func NewBalanceDb(db *sql.DB) *BalanceDb {
	return &BalanceDb{
		DB: db,
	}
}

func (b *BalanceDb) Get(accountId string) *balance.Balance {
	balance := balance.Balance{}

	stmt, _ := b.DB.Prepare("SELECT account_id, balance FROM balances WHERE account_id = ?")

	defer stmt.Close()

	row := stmt.QueryRow(accountId)

	err := row.Scan(&balance.AccountId, &balance.Amount)

	if err != nil {
		return nil
	}

	return &balance
}

func (b *BalanceDb) Create(balance balance.Balance) error {
	stmt, err := b.DB.Prepare("INSERT INTO balances (account_id, balance) VALUES (?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(balance.AccountId, balance.Amount)

	return err
}

func (b *BalanceDb) Update(balance balance.Balance) error {
	stmt, err := b.DB.Prepare("UPDATE balances SET balance = ? WHERE account_id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	fmt.Println("*** Updating balance: ", balance.Amount, balance.AccountId)

	_, err = stmt.Exec(balance.Amount, balance.AccountId)

	return err
}
