package hr_test

import (
	"testing"

	"lab2/pizzeria/hr"
)

func TestIDGeneratorAndEmailAccount(t *testing.T) {
	gen := hr.NewIDGenerator("EMP")
	id1 := gen.Next()
	gen.Reset()
	if id1 == gen.Next() {
		t.Fatalf("ids should be unique even after reset once")
	}

	email := hr.NewEmailAccount("chef@pizza.io")
	if _, err := email.Send("hello"); err == nil {
		t.Fatalf("expected inactive email error")
	}
	email.Activate()
	email.AddLabel("priority")
	if _, err := email.Send("shift"); err != nil {
		t.Fatalf("send failed: %v", err)
	}
}

func TestTasksTeamsAndHiring(t *testing.T) {
	status := hr.Status{Name: "ok", Level: 1}
	status.Update("critical", 5)
	status.Annotate("smoke")
	if !status.IsCritical() {
		t.Fatalf("status should be critical")
	}

	task := hr.EmployeeTask{Title: "Prep", Status: status, Minutes: 15}
	if err := task.MarkDone(); err != nil {
		t.Fatalf("mark done failed: %v", err)
	}
	if err := task.MarkDone(); err == nil {
		t.Fatalf("expected duplicate completion error")
	}
	if task.Describe() == "" {
		t.Fatalf("task description empty")
	}

	perm := hr.Permission{Code: "task", Scope: "kitchen"}
	perm.Allow()
	emp := hr.Employee{Name: "Ada", Email: hr.NewEmailAccount("ada@pizza.io"), Permissions: []hr.Permission{perm}}
	if err := emp.AssignTask(task); err != nil {
		t.Fatalf("assign failed: %v", err)
	}
	if emp.TaskCount() != 1 {
		t.Fatalf("task count mismatch")
	}
	if !emp.HasPermission("task") {
		t.Fatalf("permission not detected")
	}
	emp.GrantPermission(hr.Permission{Code: "extra"})

	team := hr.Team{Name: "Kitchen", Capacity: 1}
	if err := team.AddMember(emp); err != nil {
		t.Fatalf("add member failed: %v", err)
	}
	if err := team.AddMember(emp); err == nil {
		t.Fatalf("expected capacity error")
	}
	if team.Utilization() <= 0 {
		t.Fatalf("team utilization should be positive")
	}

	hrDept := hr.HRDepartment{Generator: hr.NewIDGenerator("HR")}
	hired := hrDept.Hire("Bob", "Chef", hr.NewEmailAccount("bob@pizza.io"))
	if hired.ID == "" {
		t.Fatalf("hire did not assign id")
	}
	emp2 := hr.Employee{Name: "X", Permissions: []hr.Permission{{Code: "audit", Allowed: true}}}
	hrDept.Employees = append(hrDept.Employees, emp2)
	if hrDept.RevokePermission("audit") == 0 {
		t.Fatalf("expected revoked permission")
	}
	hrDept.RevokePermission("task")
}
