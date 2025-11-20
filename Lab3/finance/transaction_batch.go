package finance

type TransactionBatch struct {
	ID        string
	Entries   []LedgerEntry
	Total     float64
	Processed bool
}

func (b *TransactionBatch) AddEntry(entry LedgerEntry) {
	b.Entries = append(b.Entries, entry)
}

func (b *TransactionBatch) ComputeTotal() {
	total := 0.0
	for _, e := range b.Entries {
		total += e.BalanceImpact()
	}
	b.Total = total
}

func (b *TransactionBatch) Close() {
	b.Processed = true
}

func (b TransactionBatch) IsBalanced() bool {
	return b.Total == 0
}
