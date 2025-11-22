package timeoff

const statusRequested = "requested"
const statusApproved = "approved"
const statusRejected = "rejected"

type LeaveRequest struct {
	ID         string
	EmployeeID string
	Days       float64
	Policy     LeavePolicy
	Status     string
}

func NewLeaveRequest(id string, employeeID string, days float64, policy LeavePolicy) LeaveRequest {
	return LeaveRequest{ID: id, EmployeeID: employeeID, Days: days, Policy: policy, Status: statusRequested}
}

func (r *LeaveRequest) Approve() {
	r.Status = statusApproved
}

func (r *LeaveRequest) Reject() {
	r.Status = statusRejected
}

func (r LeaveRequest) IsApproved() bool {
	return r.Status == statusApproved
}
