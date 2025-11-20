package pizzeria

import "lab2/pizzeria/hr"

type (
	IDGenerator  = hr.IDGenerator
	EmailAccount = hr.EmailAccount
	Status       = hr.Status
	EmployeeTask = hr.EmployeeTask
	Permission   = hr.Permission
	Employee     = hr.Employee
	Team         = hr.Team
	HRDepartment = hr.HRDepartment
)

var (
	NewIDGenerator  = hr.NewIDGenerator
	NewEmailAccount = hr.NewEmailAccount
)
