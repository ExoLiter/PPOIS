package finance

import (
	"fmt"

	"lab2/pizzeria/errors"
)

type Transaction struct {
	ID        string
	From      *BankAccount
	To        *BankAccount
	Amount    float64
	Converter CurrencyConverter
}

func (t *Transaction) Process() error {
	if t.From == nil || t.To == nil {
		return errors.TransactionDeclinedError{Reason: "missing accounts"}
	}
	if err := t.From.Withdraw(t.Amount); err != nil {
		return err
	}
	converted, err := t.Converter.Convert(t.Amount, t.To.Currency)
	if err != nil {
		return err
	}
	t.To.Deposit(converted)
	return nil
}

func (t Transaction) Describe() string {
	return fmt.Sprintf("tx:%s:%.2f", t.ID, t.Amount)
}
