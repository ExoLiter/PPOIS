package payroll

import (
	"testing"

	"lab3/hr"
)

func TestSalaryBand(t *testing.T) {
	band := SalaryBand{Level: 1, Min: 1000, Max: 2000, Currency: "USD"}
	if !band.Contains(1500) {
		t.Fatalf("amount not contained")
	}
	band.Adjust(100)
	if band.Min != 1100 || band.Max != 2250 {
		t.Fatalf("band adjust failed")
	}
}

func TestPayComponent(t *testing.T) {
	component := PayComponent{Name: "bonus", Amount: 100, Taxable: false, Recurring: true}
	if !component.IsBonus() {
		t.Fatalf("should be bonus")
	}
	component.MakeNonRecurring()
	if component.Recurring {
		t.Fatalf("should be non recurring")
	}
	if component.ApplyTo(900) != 1000 {
		t.Fatalf("apply failed")
	}
}

func TestPayCalendar(t *testing.T) {
	calendar := PayCalendar{Frequency: "monthly", CycleDays: 30, NextCutoff: "2025-01-15"}
	if !calendar.IsValid() {
		t.Fatalf("calendar should be valid")
	}
	calendar.ShiftCutoff("2025-02-15")
	if calendar.NextCutoff != "2025-02-15" {
		t.Fatalf("cutoff not shifted")
	}
}

func TestDeduction(t *testing.T) {
	d := Deduction{Name: "tax", Rate: 0.2, Cap: 100}
	if d.Calculate(1000) != 100 {
		t.Fatalf("cap not applied")
	}
	if d.RateApplied() != 0.2 {
		t.Fatalf("rate mismatch")
	}
}

func TestBonus(t *testing.T) {
	bonus := Bonus{Title: "spot", Amount: 50, Recurring: true}
	if bonus.Apply(100) != 150 {
		t.Fatalf("bonus apply failed")
	}
	bonus.StopRecurring()
	if bonus.Recurring {
		t.Fatalf("bonus should stop recurring")
	}
	if !bonus.IsActive() {
		t.Fatalf("bonus should be active")
	}
}

func TestOvertimePolicy(t *testing.T) {
	policy := OvertimePolicy{Name: "standard", RatePerHour: 10, MaxHours: 5}
	if policy.Calculate(10) != 75 {
		t.Fatalf("calculation incorrect")
	}
	policy.EnableApproval()
	if !policy.RequiresApproval {
		t.Fatalf("approval not enabled")
	}
	if !policy.Allows(4) {
		t.Fatalf("should allow hours")
	}
}

func TestPaycheck(t *testing.T) {
	paycheck := Paycheck{EmployeeID: "E1", Gross: 1000}
	paycheck.AddBonus(PayComponent{Name: "bonus", Amount: 100})
	paycheck.AddDeduction(PayComponent{Name: "tax", Amount: 50})
	paycheck.ComputeNet()
	if paycheck.Net != 1050 {
		t.Fatalf("net calculation failed")
	}
	if paycheck.TotalAdjustments() != 50 {
		t.Fatalf("adjustments mismatch")
	}
}

func TestPayrollRun(t *testing.T) {
	run := PayrollRun{ID: "R1", Period: "2025-01", Calendar: PayCalendar{Frequency: "monthly", CycleDays: 30, NextCutoff: "2025-01-15"}}
	run.AddComponent(PayComponent{Name: "bonus", Amount: 50})
	run.AddEmployee(hr.Employee{ID: "E1", Salary: 1000})
	run.AddEmployee(hr.Employee{ID: "E2", Salary: 500})
	if !run.UsesCalendar() {
		t.Fatalf("calendar not considered")
	}
	if run.TotalPayroll() != 1600 {
		t.Fatalf("total payroll mismatch")
	}
}
