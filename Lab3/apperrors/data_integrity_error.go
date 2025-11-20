package apperrors

const dataIntegrityCode = "DATA_INTEGRITY"

type DataIntegrityError struct {
	Field string
	Value string
}

func (e DataIntegrityError) Error() string {
	return dataIntegrityCode + ":" + e.Field
}

func (e DataIntegrityError) Code() string {
	return dataIntegrityCode
}

func (e DataIntegrityError) HasValue() bool {
	return e.Value != ""
}
