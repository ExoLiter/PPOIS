package hr

type Schedule struct {
	EmployeeID  string
	WeeklyHours float64
	Flexible    bool
	Shifts      []string
}

func (s *Schedule) AddShift(shift string, hours float64) {
	s.Shifts = append(s.Shifts, shift)
	s.WeeklyHours += hours
}

func (s *Schedule) ToggleFlexible() {
	s.Flexible = !s.Flexible
}

func (s Schedule) HasShift(name string) bool {
	for _, shift := range s.Shifts {
		if shift == name {
			return true
		}
	}
	return false
}
