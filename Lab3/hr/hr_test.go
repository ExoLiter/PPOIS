package hr

import "testing"

func TestEmployeeLifecycle(t *testing.T) {
	position := Position{Title: "Engineer", Level: 1, BaseSalary: 1000}
	contract := Contract{ID: "C1", EmployeeID: "E1", StartDate: "2025-01-01"}
	employee := Employee{ID: "E1", Name: "Alice", Position: position, Contract: contract, Salary: 1200, Department: "Tech"}
	employee.Activate()
	if !employee.Active {
		t.Fatalf("employee should be active")
	}
	employee.UpdateSalary(-1500)
	if employee.Salary != 0 {
		t.Fatalf("salary should not be negative")
	}
	employee.AssignPosition(Position{Title: "Senior", Level: 3, BaseSalary: 2000})
	if !employee.Position.IsSenior() {
		t.Fatalf("position should be senior")
	}
	employee.AwardBonus()
	if employee.Salary != bonusIncrease {
		t.Fatalf("bonus not applied")
	}
}

func TestPositionPromotion(t *testing.T) {
	position := Position{Title: "Specialist", Level: 2, BaseSalary: 1500}
	position.Promote()
	if position.Level != 3 {
		t.Fatalf("level not incremented")
	}
	if position.BaseSalary != 1550 {
		t.Fatalf("base salary not increased")
	}
	if !position.IsSenior() {
		t.Fatalf("senior expected")
	}
}

func TestContract(t *testing.T) {
	contract := Contract{ID: "C2", EmployeeID: "E2", StartDate: "2025-01-01", EndDate: "2025-12-31"}
	if !contract.IsActive("2025-05-01") {
		t.Fatalf("contract should be active")
	}
	contract.Extend("2026-12-31")
	if contract.EndDate != "2026-12-31" {
		t.Fatalf("contract not extended")
	}
}

func TestPermission(t *testing.T) {
	permission := Permission{Name: "payroll", Scope: "finance"}
	permission.Grant()
	if !permission.CanAccess("finance") {
		t.Fatalf("access should be granted")
	}
	permission.Revoke()
	if permission.Allowed {
		t.Fatalf("access should be revoked")
	}
}

func TestAccessBadge(t *testing.T) {
	badge := AccessBadge{ID: "B1", EmployeeID: "E3"}
	if badge.UseBadge("HQ") != "" {
		t.Fatalf("inactive badge should not work")
	}
	badge.ActivateBadge()
	location := badge.UseBadge("HQ")
	if location != "HQ" || badge.LastUsed != "HQ" {
		t.Fatalf("badge use not recorded")
	}
	if !badge.MatchesUser("E3") {
		t.Fatalf("user mismatch")
	}
}

func TestSchedule(t *testing.T) {
	schedule := Schedule{EmployeeID: "E4"}
	schedule.AddShift("morning", 4)
	schedule.AddShift("evening", 4)
	if schedule.WeeklyHours != 8 {
		t.Fatalf("hours mismatch")
	}
	if !schedule.HasShift("morning") {
		t.Fatalf("missing shift")
	}
	flexible := schedule.Flexible
	schedule.ToggleFlexible()
	if flexible == schedule.Flexible {
		t.Fatalf("flexible flag not toggled")
	}
}
