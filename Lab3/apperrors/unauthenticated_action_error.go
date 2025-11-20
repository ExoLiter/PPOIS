package apperrors

const unauthenticatedActionCode = "UNAUTHENTICATED"

type UnauthenticatedActionError struct {
	Action string
}

func (e UnauthenticatedActionError) Error() string {
	return unauthenticatedActionCode + ":" + e.Action
}

func (e UnauthenticatedActionError) Code() string {
	return unauthenticatedActionCode
}

func (e UnauthenticatedActionError) IsPublic() bool {
	return e.Action == "status"
}
