package apperrors

const approvalMissingCode = "APPROVAL_MISSING"

type ApprovalMissingError struct {
	RequestID string
	Level     int
}

func (e ApprovalMissingError) Error() string {
	return approvalMissingCode + ":" + e.RequestID
}

func (e ApprovalMissingError) Code() string {
	return approvalMissingCode
}

func (e ApprovalMissingError) NeedsManager() bool {
	return e.Level > 0
}
