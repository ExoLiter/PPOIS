package errors

import "fmt"

type BudgetExceededError struct {
	Category string
}

func (e BudgetExceededError) Error() string {
	return fmt.Sprintf("budget exceeded for %s", e.Category)
}
