package timeoff

import "testing"

func TestLeavePolicy(t *testing.T) {
	policy := LeavePolicy{Name: "Standard", AnnualLimit: 20}
	if !policy.Allow(10) {
		t.Fatalf("should allow days")
	}
	policy.EnableCarryOver()
	if !policy.CarryOver {
		t.Fatalf("carry over not enabled")
	}
	policy.SetApproval(true)
	if !policy.RequiresApproval {
		t.Fatalf("approval should be required")
	}
}

func TestLeaveBalance(t *testing.T) {
	policy := LeavePolicy{Name: "Standard", AnnualLimit: 20}
	balance := LeaveBalance{EmployeeID: "E1", Policy: policy, Available: 5}
	balance.Accrue(3)
	if balance.Available != 8 {
		t.Fatalf("accrual mismatch")
	}
	if !balance.Take(4) {
		t.Fatalf("take should succeed")
	}
	if balance.Remaining() != 4 {
		t.Fatalf("remaining mismatch")
	}
	if balance.Take(10) {
		t.Fatalf("should not take more than available")
	}
}

func TestLeaveRequestAndApproval(t *testing.T) {
	policy := LeavePolicy{Name: "Standard", AnnualLimit: 20}
	request := NewLeaveRequest("R1", "E1", 3, policy)
	if request.Status != statusRequested {
		t.Fatalf("status mismatch")
	}
	approval := LeaveApproval{Approver: "manager", Request: request}
	approval.SignOff(true, "ok")
	if !approval.Approved || approval.Request.Status != statusApproved {
		t.Fatalf("approval not recorded")
	}
	if !approval.IsFinal() {
		t.Fatalf("should be final")
	}
	approval.SignOff(false, "deny")
	if approval.Request.Status != statusRejected {
		t.Fatalf("status should change after rejection")
	}
}

func TestLeaveAccrual(t *testing.T) {
	accrual := LeaveAccrual{EmployeeID: "E2", Rate: 1.5}
	accrual.ApplyMonth()
	if accrual.Accumulated != 1.5 {
		t.Fatalf("accrual not applied")
	}
	balance := LeaveBalance{EmployeeID: "E2", Available: 2}
	if accrual.NextBalance(balance) != 3.5 {
		t.Fatalf("next balance mismatch")
	}
	accrual.Reset()
	if accrual.Accumulated != 0 {
		t.Fatalf("reset failed")
	}
}
