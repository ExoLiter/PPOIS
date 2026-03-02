package payroll

type Paycheck struct {
	EmployeeID string
	Gross      float64
	Net        float64
	Deductions []PayComponent
	Bonuses    []PayComponent
}

func (p *Paycheck) AddDeduction(d PayComponent) {
	p.Deductions = append(p.Deductions, d)
}

func (p *Paycheck) AddBonus(b PayComponent) {
	p.Bonuses = append(p.Bonuses, b)
}

func (p *Paycheck) ComputeNet() {
	total := p.Gross
	for _, bonus := range p.Bonuses {
		total += bonus.Amount
	}
	for _, deduction := range p.Deductions {
		total -= deduction.Amount
	}
	if total < 0 {
		total = 0
	}
	p.Net = total
}

func (p Paycheck) TotalAdjustments() float64 {
	total := 0.0
	for _, bonus := range p.Bonuses {
		total += bonus.Amount
	}
	for _, deduction := range p.Deductions {
		total -= deduction.Amount
	}
	return total
}
