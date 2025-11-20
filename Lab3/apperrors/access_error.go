package apperrors

const accessErrorCode = "ACCESS_DENIED"

type AccessError struct {
	Action string
	User   string
}

func (e AccessError) Error() string {
	return accessErrorCode + ":" + e.Action + ":" + e.User
}

func (e AccessError) Code() string {
	return accessErrorCode
}

func (e AccessError) IsCritical() bool {
	return e.Action == "" || e.User == ""
}
