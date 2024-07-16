package account

import "bankApp/pkg/repository"

type BankAccount interface {
	CreateUser() (int, error)
	Deposit(userId int, amount float64) error
	Withdraw(userId int, amount float64) error
	GetBalance(userId int) (float64, error)
}

type Account struct {
	BankAccount
}

func NewAccount(repos *repository.Repository) *Account {
	return &Account{
		BankAccount: NewBankAppAccount(repos.BankAccount),
	}
}
