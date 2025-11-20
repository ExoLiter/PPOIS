package performance

import (
	"testing"

	"lab3/hr"
)

func TestPerformanceGoal(t *testing.T) {
	goal := NewPerformanceGoal("G1", "E1", 100)
	goal.MarkProgress(40)
	if goal.Status != goalInProgress {
		t.Fatalf("status mismatch")
	}
	goal.MarkProgress(60)
	if goal.Status != goalCompleted {
		t.Fatalf("goal should be completed")
	}
	if goal.Completion() != 1 {
		t.Fatalf("completion ratio incorrect")
	}
}

func TestPerformanceReview(t *testing.T) {
	review := PerformanceReview{ID: "R1", EmployeeID: "E1", Score: 2.5}
	if review.IsPassing() {
		t.Fatalf("should not pass")
	}
	review.UpdateScore(3.5)
	review.AddFeedback("good job")
	if !review.IsPassing() {
		t.Fatalf("should pass")
	}
	if review.Feedback == "" {
		t.Fatalf("feedback not set")
	}
}

func TestPromotionCase(t *testing.T) {
	candidate := hr.Employee{ID: "E1", Position: hr.Position{Title: "Engineer", Level: 1, BaseSalary: 100}}
	proposed := hr.Position{Title: "Senior", Level: 2, BaseSalary: 200}
	caseItem := PromotionCase{ID: "P1", Candidate: candidate, ProposedRole: proposed}
	caseItem.Approve()
	if !caseItem.Approved || caseItem.Outcome() != "approved" {
		t.Fatalf("case not approved")
	}
	if caseItem.Candidate.Position.Title != "Senior" {
		t.Fatalf("position not updated")
	}
	caseItem.Deny("budget")
	if caseItem.Outcome() != "budget" {
		t.Fatalf("deny reason not recorded")
	}
}

func TestTrainingCourse(t *testing.T) {
	course := TrainingCourse{Name: "Safety", Hours: 4}
	course.Complete(0.6)
	if !course.Completed {
		t.Fatalf("course should be completed")
	}
	if !course.NeedsRetake() {
		t.Fatalf("should need retake")
	}
	if course.DurationHours() != 4 {
		t.Fatalf("duration mismatch")
	}
}

func TestCertification(t *testing.T) {
	cert := Certification{Name: "PMP", Issuer: "PMI", ValidUntil: "2025-12-31"}
	cert.Activate()
	if !cert.IsValid("2025-01-01") {
		t.Fatalf("cert should be valid")
	}
	cert.Expire("2024-12-31")
	if cert.Active {
		t.Fatalf("cert should be inactive")
	}
}
