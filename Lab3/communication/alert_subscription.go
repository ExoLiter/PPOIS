package communication

type AlertSubscription struct {
	ID         string
	EmployeeID string
	Topic      string
	Active     bool
}

func (s *AlertSubscription) Subscribe() {
	s.Active = true
}

func (s *AlertSubscription) Unsubscribe() {
	s.Active = false
}

func (s AlertSubscription) IsActive() bool {
	return s.Active
}
