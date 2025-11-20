package communication

type NotificationPreference struct {
	EmployeeID string
	Channel    string
	Enabled    bool
	Frequency  string
}

func (p *NotificationPreference) Enable() {
	p.Enabled = true
}

func (p *NotificationPreference) Disable() {
	p.Enabled = false
}

func (p *NotificationPreference) UpdateFrequency(frequency string) {
	p.Frequency = frequency
}
