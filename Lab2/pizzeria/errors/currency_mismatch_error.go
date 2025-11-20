package errors

import "fmt"

type CurrencyMismatchError struct {
	From string
	To   string
}

func (e CurrencyMismatchError) Error() string {
	return fmt.Sprintf("currency mismatch: %s to %s", e.From, e.To)
}
