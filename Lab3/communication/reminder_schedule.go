package communication

type ReminderSchedule struct {
	ID       string
	Times    []string
	Enabled  bool
	LastSent string
}

func (s *ReminderSchedule) AddTime(time string) {
	s.Times = append(s.Times, time)
}

func (s *ReminderSchedule) Toggle() {
	s.Enabled = !s.Enabled
}

func (s *ReminderSchedule) RecordSend(time string) {
	s.LastSent = time
}

func (s ReminderSchedule) NextTime() string {
	if len(s.Times) == 0 {
		return ""
	}
	return s.Times[0]
}
