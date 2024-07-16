package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BankAppPostgres struct {
	db *sqlx.DB
}

func NewBankAppPostgres(db *sqlx.DB) *BankAppPostgres {
	return &BankAppPostgres{db: db}
}

func (b *BankAppPostgres) CreateUser() (int, error) {
	var userId int

	query := fmt.Sprintf("INSERT INTO %s DEFAULT VALUES RETURNING id", usersTable)
	row := b.db.QueryRow(query)

	if err := row.Scan(&userId); err != nil {
		return 0, err
	}

	return userId, nil
}

func (b *BankAppPostgres) Deposit(userId int, amount float64) error {
	query := fmt.Sprintf("UPDATE %s SET balance = balance + $1 WHERE id = $2", usersTable)
	result, err := b.db.Exec(query, amount, userId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", userId)
	}

	return nil
}

func (b *BankAppPostgres) Withdraw(userId int, amount float64) error {
	query := fmt.Sprintf("UPDATE %s SET balance = balance - $1 WHERE id = $2", usersTable)
	result, err := b.db.Exec(query, amount, userId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", userId)
	}

	return nil
}

func (b *BankAppPostgres) GetBalance(userId int) (float64, error) {
	var balance float64
	query := fmt.Sprintf("SELECT balance FROM %s WHERE id = $1", usersTable)
	err := b.db.Get(&balance, query, userId)

	if err != nil {
		return 0, err
	}
	return balance, err
}
