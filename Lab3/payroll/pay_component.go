package payroll

type PayComponent struct {
	Name      string
	Amount    float64
	Taxable   bool
	Recurring bool
}

func (c PayComponent) ApplyTo(base float64) float64 {
	return base + c.Amount
}

func (c *PayComponent) MakeNonRecurring() {
	c.Recurring = false
}

func (c PayComponent) IsBonus() bool {
	return c.Amount > 0 && !c.Taxable
}
