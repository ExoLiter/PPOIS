package apperrors

const inconsistentStateCode = "INCONSISTENT_STATE"

type InconsistentStateError struct {
	Entity string
	Detail string
}

func (e InconsistentStateError) Error() string {
	return inconsistentStateCode + ":" + e.Entity + ":" + e.Detail
}

func (e InconsistentStateError) Code() string {
	return inconsistentStateCode
}

func (e InconsistentStateError) IsEmpty() bool {
	return e.Entity == "" && e.Detail == ""
}
