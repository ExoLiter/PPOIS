package errors

import "fmt"

type OrderValidationError struct {
	Order string
}

func (e OrderValidationError) Error() string {
	return fmt.Sprintf("order %s is invalid", e.Order)
}
