package benefits

const reimbursementPending = "pending"
const reimbursementApproved = "approved"
const reimbursementRejected = "rejected"

type ReimbursementRequest struct {
	ID         string
	EmployeeID string
	Amount     float64
	Category   string
	Status     string
}

func NewReimbursementRequest(id string, employeeID string, amount float64, category string) ReimbursementRequest {
	return ReimbursementRequest{ID: id, EmployeeID: employeeID, Amount: amount, Category: category, Status: reimbursementPending}
}

func (r *ReimbursementRequest) Approve() {
	r.Status = reimbursementApproved
}

func (r *ReimbursementRequest) Reject() {
	r.Status = reimbursementRejected
}

func (r ReimbursementRequest) IsApproved() bool {
	return r.Status == reimbursementApproved
}
