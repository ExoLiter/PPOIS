package finance

import "testing"

func TestCurrencyFormat(t *testing.T) {
	currency := Currency{Code: "USD", Symbol: "$", Precision: 2}
	if currency.Format(10.257) != "$10.26" {
		t.Fatalf("format mismatch")
	}
	if !currency.Equals("USD") {
		t.Fatalf("code mismatch")
	}
}

func TestExchangeRateTable(t *testing.T) {
	table := NewExchangeRateTable(Currency{Code: "USD"})
	table.SetRate("EUR", 0.9)
	if !table.HasRate("EUR") {
		t.Fatalf("rate missing")
	}
	if table.Convert(100, "EUR") != 90 {
		t.Fatalf("conversion mismatch")
	}
	if table.Convert(10, "GBP") != 0 {
		t.Fatalf("unknown rate should be zero")
	}
}

func TestBankAccount(t *testing.T) {
	account := BankAccount{Number: "123", Owner: "Alice", Currency: Currency{Code: "USD"}}
	account.Activate()
	account.Deposit(200)
	if !account.Withdraw(50) {
		t.Fatalf("withdraw should succeed")
	}
	if account.Balance != 150 {
		t.Fatalf("balance mismatch")
	}
	if account.Withdraw(200) {
		t.Fatalf("overdraft should fail")
	}
	if !account.HasCurrency("USD") {
		t.Fatalf("currency mismatch")
	}
}

func TestPaymentInstruction(t *testing.T) {
	instruction := NewPaymentInstruction("P1", 100, Currency{Code: "USD"}, BankAccount{Number: "321"})
	if instruction.Status != statusPending {
		t.Fatalf("pending expected")
	}
	instruction.Approve()
	if !instruction.IsReady() {
		t.Fatalf("should be ready")
	}
	instruction.MarkPaid()
	if instruction.Status != statusPaid {
		t.Fatalf("should be paid")
	}
}

func TestLedgerEntryAndBatch(t *testing.T) {
	account := BankAccount{Number: "999"}
	entry1 := LedgerEntry{ID: "L1", Debit: 100, Credit: 0, Account: account}
	entry2 := LedgerEntry{ID: "L2", Debit: 0, Credit: 100, Account: account}
	if entry1.BalanceImpact() != -100 {
		t.Fatalf("impact mismatch")
	}
	if !entry2.IsCredit() {
		t.Fatalf("should be credit")
	}
	batch := TransactionBatch{ID: "B1"}
	batch.AddEntry(entry1)
	batch.AddEntry(entry2)
	batch.ComputeTotal()
	if batch.Total != 0 {
		t.Fatalf("total should balance")
	}
	if !batch.IsBalanced() {
		t.Fatalf("batch should be balanced")
	}
	batch.Close()
	if !batch.Processed {
		t.Fatalf("batch should be processed")
	}
}
