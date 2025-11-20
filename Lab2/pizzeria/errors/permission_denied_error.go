package errors

import "fmt"

type PermissionDeniedError struct {
	Code string
}

func (e PermissionDeniedError) Error() string {
	return fmt.Sprintf("permission %s denied", e.Code)
}
