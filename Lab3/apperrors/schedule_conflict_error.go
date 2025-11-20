package apperrors

const scheduleConflictCode = "SCHEDULE_CONFLICT"

type ScheduleConflictError struct {
	EmployeeID string
	Slot       string
}

func (e ScheduleConflictError) Error() string {
	return scheduleConflictCode + ":" + e.EmployeeID + ":" + e.Slot
}

func (e ScheduleConflictError) Code() string {
	return scheduleConflictCode
}

func (e ScheduleConflictError) IsSameSlot(slot string) bool {
	return e.Slot == slot
}
