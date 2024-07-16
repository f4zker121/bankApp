package repository

import (
	"github.com/jmoiron/sqlx"
)

type BankAccount interface {
	CreateUser() (int, error)
	Deposit(userId int, amount float64) error
	Withdraw(userId int, amount float64) error
	GetBalance(userId int) (float64, error)
}

type Repository struct {
	BankAccount
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		BankAccount: NewBankAppPostgres(db),
	}
}
