package timeoff

type LeaveBalance struct {
	EmployeeID string
	Available  float64
	Taken      float64
	Policy     LeavePolicy
}

func (b *LeaveBalance) Accrue(days float64) {
	b.Available += days
}

func (b *LeaveBalance) Take(days float64) bool {
	if days > b.Available {
		return false
	}
	b.Available -= days
	b.Taken += days
	return true
}

func (b LeaveBalance) Remaining() float64 {
	return b.Available
}
