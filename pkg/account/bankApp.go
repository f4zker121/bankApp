package account

import (
	"bankApp/pkg/repository"
)

type BankAppAccount struct {
	repo repository.BankAccount
}

func NewBankAppAccount(repo repository.BankAccount) *BankAppAccount {
	return &BankAppAccount{repo: repo}
}

func (b *BankAppAccount) CreateUser() (int, error) {
	return b.repo.CreateUser()
}

func (b *BankAppAccount) Deposit(userId int, amount float64) error {
	return b.repo.Deposit(userId, amount)
}

func (b *BankAppAccount) Withdraw(userId int, amount float64) error {
	return b.repo.Withdraw(userId, amount)
}

func (b *BankAppAccount) GetBalance(userId int) (float64, error) {
	return b.repo.GetBalance(userId)
}
