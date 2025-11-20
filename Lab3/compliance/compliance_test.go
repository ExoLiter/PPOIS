package compliance

import "testing"

func TestPolicyDocument(t *testing.T) {
	doc := PolicyDocument{Name: "Security"}
	doc.Publish("v1")
	if !doc.IsActive() || doc.Version != "v1" {
		t.Fatalf("publish failed")
	}
	doc.Retire()
	if doc.Active {
		t.Fatalf("document should be retired")
	}
}

func TestAuditLog(t *testing.T) {
	log := AuditLog{Actor: "alice", Action: "login", Timestamp: "2025-01-01"}
	log.AddEntry("first")
	log.AddEntry("second")
	if log.EntryCount() != 2 {
		t.Fatalf("entry count mismatch")
	}
	if log.LastEntry() != "second" {
		t.Fatalf("last entry mismatch")
	}
}

func TestAccessReview(t *testing.T) {
	review := AccessReview{ID: "R1", Reviewer: "bob"}
	if review.HasFindings() {
		t.Fatalf("should start empty")
	}
	review.AddFinding("missing mfa")
	if !review.HasFindings() {
		t.Fatalf("findings not recorded")
	}
	review.Approve()
	if !review.Approved {
		t.Fatalf("review should be approved")
	}
}

func TestIncidentReport(t *testing.T) {
	report := IncidentReport{ID: "I1", Severity: severityHigh, Description: "data loss"}
	if !report.IsHighSeverity() {
		t.Fatalf("severity check failed")
	}
	report.Resolve("restored backup")
	if !report.Resolved || report.Resolution == "" {
		t.Fatalf("resolution not recorded")
	}
}

func TestRiskAssessment(t *testing.T) {
	assessment := NewRiskAssessment("RA1", 5)
	assessment.AddRisk("auth", 3)
	assessment.AddRisk("payroll", 4)
	if !assessment.IsCritical() {
		t.Fatalf("should be critical")
	}
	if assessment.Score != 7 {
		t.Fatalf("score mismatch")
	}
}
