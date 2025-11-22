package apperrors

const payrollLockCode = "PAYROLL_LOCKED"

type PayrollLockError struct {
	Period string
	Owner  string
}

func (e PayrollLockError) Error() string {
	return payrollLockCode + ":" + e.Period + ":" + e.Owner
}

func (e PayrollLockError) Code() string {
	return payrollLockCode
}

func (e PayrollLockError) RequiresEscalation() bool {
	return e.Owner != ""
}
