package apperrors

const insufficientBalanceCode = "INSUFFICIENT_BALANCE"

type InsufficientBalanceError struct {
	Account string
	Needed  float64
	Actual  float64
}

func (e InsufficientBalanceError) Error() string {
	return insufficientBalanceCode + ":" + e.Account
}

func (e InsufficientBalanceError) Code() string {
	return insufficientBalanceCode
}

func (e InsufficientBalanceError) Shortfall() float64 {
	if e.Needed < e.Actual {
		return 0
	}
	return e.Needed - e.Actual
}
