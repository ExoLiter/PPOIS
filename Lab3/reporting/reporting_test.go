package reporting

import "testing"

func TestPayrollReport(t *testing.T) {
	report := PayrollReport{Period: "2025-01", GeneratedBy: "system"}
	report.AddItem(100)
	report.AddItem(200)
	report.ComputeTotal()
	if !report.HasData() {
		t.Fatalf("report should have data")
	}
	if report.Total != 300 {
		t.Fatalf("total mismatch")
	}
}

func TestHeadcountReport(t *testing.T) {
	report := NewHeadcountReport()
	report.AddDepartment("Tech", 5)
	report.AddDepartment("HR", 2)
	if report.Total() != 7 {
		t.Fatalf("total mismatch")
	}
	if report.DepartmentCount("Tech") != 5 {
		t.Fatalf("department count mismatch")
	}
}

func TestComplianceReport(t *testing.T) {
	report := ComplianceReport{Reviewer: "alice"}
	if !report.IsClean() {
		t.Fatalf("report should be clean")
	}
	report.AddIssue("missing logs", 2)
	if report.IsClean() {
		t.Fatalf("report should not be clean")
	}
	if report.RiskLevel != 2 {
		t.Fatalf("risk mismatch")
	}
}

func TestBenefitsReport(t *testing.T) {
	report := BenefitsReport{EmployeeID: "E1"}
	report.AddBenefit(100)
	report.AddBenefit(50)
	if report.AverageCost() != 75 {
		t.Fatalf("average cost mismatch")
	}
	if report.PlanCount != 2 {
		t.Fatalf("plan count mismatch")
	}
}
