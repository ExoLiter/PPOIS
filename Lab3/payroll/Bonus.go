package payroll

type Bonus struct {
	Title     string
	Amount    float64
	Recurring bool
}

func (b Bonus) Apply(base float64) float64 {
	return base + b.Amount
}

func (b *Bonus) StopRecurring() {
	b.Recurring = false
}

func (b Bonus) IsActive() bool {
	return b.Amount > 0
}
