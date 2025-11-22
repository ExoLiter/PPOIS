package payroll

const overtimeMultiplier = 1.5

type OvertimePolicy struct {
	Name          string
	RatePerHour   float64
	MaxHours      float64
	RequiresApproval bool
}

func (p OvertimePolicy) Calculate(hours float64) float64 {
	if hours > p.MaxHours {
		hours = p.MaxHours
	}
	return hours * p.RatePerHour * overtimeMultiplier
}

func (p *OvertimePolicy) EnableApproval() {
	p.RequiresApproval = true
}

func (p OvertimePolicy) Allows(hours float64) bool {
	return hours <= p.MaxHours
}
