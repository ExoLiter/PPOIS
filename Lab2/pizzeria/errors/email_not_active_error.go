package errors

import "fmt"

type EmailNotActiveError struct {
	Address string
}

func (e EmailNotActiveError) Error() string {
	return fmt.Sprintf("email %s is not active", e.Address)
}
