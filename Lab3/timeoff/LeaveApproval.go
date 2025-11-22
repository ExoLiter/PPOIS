package timeoff

type LeaveApproval struct {
	Approver string
	Request  LeaveRequest
	Comment  string
	Approved bool
}

func (a *LeaveApproval) SignOff(approve bool, comment string) {
	a.Approved = approve
	a.Comment = comment
	if approve {
		a.Request.Approve()
	} else {
		a.Request.Reject()
	}
}

func (a LeaveApproval) IsFinal() bool {
	return a.Request.Status == statusApproved || a.Request.Status == statusRejected
}
