package timeoff

const approvalRequired = "pending"

type LeavePolicy struct {
	Name              string
	AnnualLimit       float64
	CarryOver         bool
	RequiresApproval  bool
}

func (p *LeavePolicy) EnableCarryOver() {
	p.CarryOver = true
}

func (p *LeavePolicy) SetApproval(required bool) {
	p.RequiresApproval = required
}

func (p LeavePolicy) Allow(days float64) bool {
	return days <= p.AnnualLimit
}
