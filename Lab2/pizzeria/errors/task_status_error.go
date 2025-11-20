package errors

import "fmt"

type TaskStatusError struct {
	Task string
}

func (e TaskStatusError) Error() string {
	return fmt.Sprintf("task %s is already completed", e.Task)
}
