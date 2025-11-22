package payroll

const minCycleDays = 7

type PayCalendar struct {
	Frequency  string
	CycleDays  int
	NextCutoff string
}

func (c PayCalendar) IsValid() bool {
	return c.CycleDays >= minCycleDays && c.NextCutoff != ""
}

func (c *PayCalendar) ShiftCutoff(newCutoff string) {
	c.NextCutoff = newCutoff
}
