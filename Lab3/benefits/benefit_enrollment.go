package benefits

const enrollmentActive = "active"
const enrollmentCancelled = "cancelled"

type BenefitEnrollment struct {
	EmployeeID string
	Plan       BenefitPlan
	Status     string
	StartDate  string
}

func NewBenefitEnrollment(employeeID string, plan BenefitPlan, startDate string) BenefitEnrollment {
	return BenefitEnrollment{EmployeeID: employeeID, Plan: plan, StartDate: startDate, Status: enrollmentActive}
}

func (e *BenefitEnrollment) Cancel() {
	e.Status = enrollmentCancelled
}

func (e *BenefitEnrollment) Reactivate() {
	e.Status = enrollmentActive
}

func (e BenefitEnrollment) IsActive() bool {
	return e.Status == enrollmentActive
}
