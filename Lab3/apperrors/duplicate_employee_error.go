package apperrors

const duplicateEmployeeCode = "DUPLICATE_EMPLOYEE"

type DuplicateEmployeeError struct {
	EmployeeID string
}

func (e DuplicateEmployeeError) Error() string {
	return duplicateEmployeeCode + ":" + e.EmployeeID
}

func (e DuplicateEmployeeError) Code() string {
	return duplicateEmployeeCode
}

func (e DuplicateEmployeeError) IsSame(id string) bool {
	return e.EmployeeID == id
}
