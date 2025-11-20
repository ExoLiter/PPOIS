package finance

type AccountingDepartment struct {
	Accounts     []BankAccount
	Transactions []Transaction
	Budget       Budget
}

func (a *AccountingDepartment) Post(tx Transaction) error {
	if err := tx.Process(); err != nil {
		return err
	}
	a.Transactions = append(a.Transactions, tx)
	return nil
}

func (a AccountingDepartment) Balance(owner string) float64 {
	total := 0.0
	for _, acc := range a.Accounts {
		if acc.Owner == owner {
			total += acc.Balance
		}
	}
	return total
}
