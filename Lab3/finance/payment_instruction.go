package finance

const statusPending = "pending"
const statusApproved = "approved"
const statusPaid = "paid"

type PaymentInstruction struct {
	ID       string
	Amount   float64
	Currency Currency
	Receiver BankAccount
	Status   string
}

func NewPaymentInstruction(id string, amount float64, currency Currency, receiver BankAccount) PaymentInstruction {
	return PaymentInstruction{ID: id, Amount: amount, Currency: currency, Receiver: receiver, Status: statusPending}
}

func (p *PaymentInstruction) Approve() {
	if p.Status == statusPending {
		p.Status = statusApproved
	}
}

func (p *PaymentInstruction) MarkPaid() {
	if p.Status == statusApproved {
		p.Status = statusPaid
	}
}

func (p PaymentInstruction) IsReady() bool {
	return p.Status == statusApproved
}
