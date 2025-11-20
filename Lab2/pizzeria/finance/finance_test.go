package finance_test

import (
	"testing"

	"lab2/pizzeria/finance"
)

func TestBudgetAndTransactions(t *testing.T) {
	budget := finance.NewBudget("Ops", 100)
	if err := budget.Allocate("utilities", 80); err != nil {
		t.Fatalf("unexpected allocation error: %v", err)
	}
	if err := budget.Allocate("extra", 30); err == nil {
		t.Fatalf("expected budget overflow")
	}

	currency := finance.Currency{Code: "USD", Symbol: "$", Rate: 1}
	currency.AdjustRate(1.5)
	converter := finance.NewCurrencyConverter(currency)
	converter.AddRate("USD", 1)

	acc1 := finance.BankAccount{Number: "1", Owner: "Ops", Balance: 100, Currency: currency}
	acc2 := finance.BankAccount{Number: "2", Owner: "Chef", Balance: 0, Currency: currency}
	if err := acc1.Withdraw(200); err == nil {
		t.Fatalf("expected insufficient funds")
	}

	tx := finance.Transaction{ID: "T1", From: &acc1, To: &acc2, Amount: 20, Converter: converter}
	if err := tx.Process(); err != nil {
		t.Fatalf("process failed: %v", err)
	}
	if desc := tx.Describe(); desc == "" {
		t.Fatalf("transaction describe empty")
	}

	accounting := finance.AccountingDepartment{Accounts: []finance.BankAccount{acc1, acc2}, Budget: budget}
	if err := accounting.Post(tx); err != nil {
		t.Fatalf("post failed: %v", err)
	}
	if accounting.Balance("Chef") <= 0 {
		t.Fatalf("balance not updated")
	}
}

func TestReceiptsAndCurrency(t *testing.T) {
	rec := finance.Receipt{Number: "R1"}
	rec.AddItem("Cheese", 3.5)
	if rec.Summary() == "" {
		t.Fatalf("receipt summary empty")
	}
	if rec.Total <= 0 {
		t.Fatalf("receipt total should accumulate")
	}

	currency := finance.Currency{Code: "USD", Symbol: "$", Rate: 1}
	if currency.Format(2) == "" {
		t.Fatalf("format empty")
	}
	converter := finance.NewCurrencyConverter(currency)
	converter.AddRate("USD", 1)
	if _, err := converter.Convert(2, currency); err != nil {
		t.Fatalf("convert failed: %v", err)
	}
	if _, err := converter.Convert(2, finance.Currency{Code: "EUR"}); err == nil {
		t.Fatalf("expected currency mismatch")
	}
}
