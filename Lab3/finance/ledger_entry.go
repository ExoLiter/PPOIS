package finance

type LedgerEntry struct {
	ID          string
	Debit       float64
	Credit      float64
	Account     BankAccount
	Description string
}

func (e LedgerEntry) BalanceImpact() float64 {
	return e.Credit - e.Debit
}

func (e LedgerEntry) IsCredit() bool {
	return e.Credit > 0
}
