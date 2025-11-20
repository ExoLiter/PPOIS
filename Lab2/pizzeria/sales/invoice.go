package sales

type Invoice struct {
	Number string
	Order  Order
	Paid   bool
}

func (i *Invoice) MarkPaid() {
	i.Paid = true
}

func (i Invoice) IsPaid() bool {
	return i.Paid
}
