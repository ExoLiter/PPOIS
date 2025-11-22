package timeoff

type LeaveAccrual struct {
	EmployeeID string
	Rate       float64
	Accumulated float64
}

func (a *LeaveAccrual) ApplyMonth() {
	a.Accumulated += a.Rate
}

func (a *LeaveAccrual) Reset() {
	a.Accumulated = 0
}

func (a LeaveAccrual) NextBalance(current LeaveBalance) float64 {
	return current.Available + a.Rate
}
