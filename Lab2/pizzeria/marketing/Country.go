package marketing

type Country struct {
	Name        string
	TaxRate     float64
	Regulations []string
}

func (c *Country) AddRegulation(rule string) {
	c.Regulations = append(c.Regulations, rule)
}

func (c Country) EffectiveTax(amount float64) float64 {
	return amount * c.TaxRate
}
