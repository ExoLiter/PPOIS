package errors

import "fmt"

type TransactionDeclinedError struct {
	Reason string
}

func (e TransactionDeclinedError) Error() string {
	return fmt.Sprintf("transaction declined: %s", e.Reason)
}
