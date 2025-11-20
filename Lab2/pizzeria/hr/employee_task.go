package hr

import (
	"fmt"
	"strings"

	"lab2/pizzeria/errors"
)

type EmployeeTask struct {
	Title   string
	Status  Status
	Minutes int
}

func (t *EmployeeTask) MarkDone() error {
	if strings.EqualFold(t.Status.Name, "done") {
		return errors.TaskStatusError{Task: t.Title}
	}
	t.Status.Update("done", 0)
	return nil
}

func (t EmployeeTask) Describe() string {
	return fmt.Sprintf("task:%s:%s", t.Title, t.Status.Name)
}
