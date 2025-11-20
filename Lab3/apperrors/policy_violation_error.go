package apperrors

const policyViolationCode = "POLICY_VIOLATION"

type PolicyViolationError struct {
	Policy string
	Actor  string
}

func (e PolicyViolationError) Error() string {
	return policyViolationCode + ":" + e.Policy + ":" + e.Actor
}

func (e PolicyViolationError) Code() string {
	return policyViolationCode
}

func (e PolicyViolationError) Involves(actor string) bool {
	return e.Actor == actor
}
