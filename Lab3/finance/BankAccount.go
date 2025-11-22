package finance

const overdraftLimit = 0.0

type BankAccount struct {
	Number   string
	Owner    string
	Balance  float64
	Currency Currency
	Active   bool
}

func (a *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		return
	}
	a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64) bool {
	if amount <= 0 {
		return false
	}
	if a.Balance-amount < overdraftLimit {
		return false
	}
	a.Balance -= amount
	return true
}

func (a *BankAccount) Activate() {
	a.Active = true
}

func (a BankAccount) HasCurrency(code string) bool {
	return a.Currency.Code == code
}
