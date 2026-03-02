package hr

import (
	"lab2/pizzeria/errors"
)

type Employee struct {
	ID          string
	Name        string
	Role        string
	Email       EmailAccount
	Permissions []Permission
	Tasks       []EmployeeTask
	Status      Status
}

func (e *Employee) AssignTask(task EmployeeTask) error {
	if !e.HasPermission("task") {
		return errors.PermissionDeniedError{Code: "task"}
	}
	e.Tasks = append(e.Tasks, task)
	return nil
}

func (e *Employee) HasPermission(code string) bool {
	for _, p := range e.Permissions {
		if p.Check(code) {
			return true
		}
	}
	return false
}

func (e *Employee) GrantPermission(permission Permission) {
	e.Permissions = append(e.Permissions, permission)
}

func (e Employee) TaskCount() int {
	return len(e.Tasks)
}
