package apperrors

import "testing"

func TestAccessError(t *testing.T) {
	err := AccessError{Action: "view", User: "alice"}
	if err.Code() != accessErrorCode {
		t.Fatalf("code mismatch")
	}
	if err.IsCritical() {
		t.Fatalf("should not be critical")
	}
	if err.Error() == "" {
		t.Fatalf("error message empty")
	}
}

func TestPayrollLockError(t *testing.T) {
	err := PayrollLockError{Period: "2025-01", Owner: "payroll"}
	if !err.RequiresEscalation() {
		t.Fatalf("should require escalation")
	}
	if err.Code() != payrollLockCode {
		t.Fatalf("code mismatch")
	}
	if err.Error() == "" {
		t.Fatalf("error message empty")
	}
}

func TestInsufficientBalanceError(t *testing.T) {
	err := InsufficientBalanceError{Account: "main", Needed: 200, Actual: 150}
	if err.Shortfall() != 50 {
		t.Fatalf("wrong shortfall")
	}
	if err.Error() == "" {
		t.Fatalf("error message empty")
	}
}

func TestDuplicateEmployeeError(t *testing.T) {
	err := DuplicateEmployeeError{EmployeeID: "E1"}
	if !err.IsSame("E1") {
		t.Fatalf("should match id")
	}
	if err.Code() != duplicateEmployeeCode {
		t.Fatalf("code mismatch")
	}
	if err.Error() == "" {
		t.Fatalf("error message empty")
	}
}

func TestInvalidCurrencyError(t *testing.T) {
	err := InvalidCurrencyError{Currency: ""}
	if !err.IsEmpty() {
		t.Fatalf("should be empty")
	}
	if err.Error() == "" {
		t.Fatalf("error message empty")
	}
	if err.Code() != invalidCurrencyCode {
		t.Fatalf("code mismatch")
	}
}

func TestScheduleConflictError(t *testing.T) {
	err := ScheduleConflictError{EmployeeID: "E2", Slot: "morning"}
	if !err.IsSameSlot("morning") {
		t.Fatalf("slot mismatch")
	}
	if err.Code() != scheduleConflictCode {
		t.Fatalf("code mismatch")
	}
	if err.Error() == "" {
		t.Fatalf("error message empty")
	}
}

func TestPolicyViolationError(t *testing.T) {
	err := PolicyViolationError{Policy: "security", Actor: "bob"}
	if !err.Involves("bob") {
		t.Fatalf("actor mismatch")
	}
	if err.Code() != policyViolationCode {
		t.Fatalf("code mismatch")
	}
	if err.Error() == "" {
		t.Fatalf("error message empty")
	}
}

func TestOvertimeLimitError(t *testing.T) {
	err := OvertimeLimitError{EmployeeID: "E3", Hours: 10, Limit: 8}
	if err.Excess() != 2 {
		t.Fatalf("excess mismatch")
	}
	if err.Code() != overtimeLimitCode {
		t.Fatalf("code mismatch")
	}
	if err.Error() == "" {
		t.Fatalf("error message empty")
	}
}

func TestUnauthenticatedActionError(t *testing.T) {
	err := UnauthenticatedActionError{Action: "status"}
	if !err.IsPublic() {
		t.Fatalf("should be public action")
	}
	if err.Error() == "" {
		t.Fatalf("error message empty")
	}
	if err.Code() != unauthenticatedActionCode {
		t.Fatalf("code mismatch")
	}
}

func TestInconsistentStateError(t *testing.T) {
	err := InconsistentStateError{}
	if !err.IsEmpty() {
		t.Fatalf("should be empty state")
	}
	if err.Code() != inconsistentStateCode {
		t.Fatalf("code mismatch")
	}
	if err.Error() == "" {
		t.Fatalf("error message empty")
	}
}

func TestApprovalMissingError(t *testing.T) {
	err := ApprovalMissingError{RequestID: "R1", Level: 1}
	if !err.NeedsManager() {
		t.Fatalf("should need manager")
	}
	if err.Error() == "" {
		t.Fatalf("error message empty")
	}
	if err.Code() != approvalMissingCode {
		t.Fatalf("code mismatch")
	}
}

func TestDataIntegrityError(t *testing.T) {
	err := DataIntegrityError{Field: "amount", Value: "10"}
	if !err.HasValue() {
		t.Fatalf("value expected")
	}
	if err.Code() != dataIntegrityCode {
		t.Fatalf("code mismatch")
	}
	if err.Error() == "" {
		t.Fatalf("error message empty")
	}
}
