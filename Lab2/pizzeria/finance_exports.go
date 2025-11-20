package pizzeria

import "lab2/pizzeria/finance"

type (
	Budget               = finance.Budget
	Currency             = finance.Currency
	BankAccount          = finance.BankAccount
	Receipt              = finance.Receipt
	CurrencyConverter    = finance.CurrencyConverter
	Transaction          = finance.Transaction
	AccountingDepartment = finance.AccountingDepartment
)

var (
	NewBudget            = finance.NewBudget
	NewCurrencyConverter = finance.NewCurrencyConverter
)
