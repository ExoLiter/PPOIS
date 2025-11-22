package apperrors

const invalidCurrencyCode = "INVALID_CURRENCY"

type InvalidCurrencyError struct {
	Currency string
}

func (e InvalidCurrencyError) Error() string {
	return invalidCurrencyCode + ":" + e.Currency
}

func (e InvalidCurrencyError) Code() string {
	return invalidCurrencyCode
}

func (e InvalidCurrencyError) IsEmpty() bool {
	return e.Currency == ""
}
