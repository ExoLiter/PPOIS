package benefits

import "testing"

func TestBenefitPlanShares(t *testing.T) {
	plan := BenefitPlan{Name: "Health", Provider: "BestCare", Cost: 200, EmployerContribution: 0.5}
	plan.Activate()
	if !plan.Active {
		t.Fatalf("plan should be active")
	}
	if plan.EmployerShare() != 100 {
		t.Fatalf("employer share mismatch")
	}
	if plan.EmployeeShare() != 100 {
		t.Fatalf("employee share mismatch")
	}
}

func TestBenefitEnrollment(t *testing.T) {
	plan := BenefitPlan{Name: "Health", Provider: "BestCare", Cost: 150, EmployerContribution: 0.4}
	enrollment := NewBenefitEnrollment("E1", plan, "2025-01-01")
	if !enrollment.IsActive() {
		t.Fatalf("enrollment should be active")
	}
	enrollment.Cancel()
	if enrollment.IsActive() {
		t.Fatalf("enrollment should be cancelled")
	}
	enrollment.Reactivate()
	if !enrollment.IsActive() {
		t.Fatalf("enrollment should reactivate")
	}
}

func TestInsuranceClaim(t *testing.T) {
	plan := BenefitPlan{Name: "Health", Provider: "BestCare", Cost: 150, EmployerContribution: 0.4}
	enrollment := NewBenefitEnrollment("E1", plan, "2025-01-01")
	claim := NewInsuranceClaim("C1", enrollment, 500)
	claim.Approve()
	if claim.Status != claimApproved || claim.Payout() != 500 {
		t.Fatalf("claim approve failed")
	}
	claim.Reject()
	if claim.Payout() != 0 || claim.Status != claimRejected {
		t.Fatalf("claim reject failed")
	}
}

func TestReimbursementRequest(t *testing.T) {
	request := NewReimbursementRequest("R1", "E2", 50, "Travel")
	if request.IsApproved() {
		t.Fatalf("should not be approved")
	}
	request.Approve()
	if !request.IsApproved() {
		t.Fatalf("should be approved")
	}
	request.Reject()
	if request.Status != reimbursementRejected {
		t.Fatalf("status should be rejected")
	}
}

func TestBenefitStatement(t *testing.T) {
	plan1 := BenefitPlan{Name: "Health", Cost: 100}
	plan2 := BenefitPlan{Name: "Dental", Cost: 50}
	statement := BenefitStatement{EmployeeID: "E3"}
	statement.AddPlan(plan1)
	statement.AddPlan(plan2)
	statement.ComputeTotal()
	if statement.TotalCost != 150 {
		t.Fatalf("total cost mismatch")
	}
	if statement.PlanCount() != 2 {
		t.Fatalf("plan count mismatch")
	}
}
