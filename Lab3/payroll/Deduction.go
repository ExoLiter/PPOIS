package payroll

type Deduction struct {
	Name string
	Rate float64
	Cap  float64
}

func (d Deduction) Calculate(amount float64) float64 {
	value := amount * d.Rate
	if value > d.Cap {
		return d.Cap
	}
	return value
}

func (d Deduction) RateApplied() float64 {
	return d.Rate
}
