package apperrors

const overtimeLimitCode = "OVERTIME_LIMIT"

type OvertimeLimitError struct {
	EmployeeID string
	Hours      float64
	Limit      float64
}

func (e OvertimeLimitError) Error() string {
	return overtimeLimitCode + ":" + e.EmployeeID
}

func (e OvertimeLimitError) Code() string {
	return overtimeLimitCode
}

func (e OvertimeLimitError) Excess() float64 {
	if e.Hours <= e.Limit {
		return 0
	}
	return e.Hours - e.Limit
}
