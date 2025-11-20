package payroll

const salaryAdjustStep = 150.0

type SalaryBand struct {
	Level    int
	Min      float64
	Max      float64
	Currency string
}

func (b *SalaryBand) Contains(amount float64) bool {
	return amount >= b.Min && amount <= b.Max
}

func (b *SalaryBand) Adjust(delta float64) {
	b.Min += delta
	b.Max += delta + salaryAdjustStep
}
