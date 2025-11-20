package finance

import "lab2/pizzeria/errors"

type BankAccount struct {
	Number   string
	Owner    string
	Balance  float64
	Currency Currency
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount > b.Balance {
		return errors.TransactionDeclinedError{Reason: "insufficient funds"}
	}
	b.Balance -= amount
	return nil
}
