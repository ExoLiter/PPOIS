package hr

type Contract struct {
	ID         string
	EmployeeID string
	StartDate  string
	EndDate    string
	Rate       float64
	Type       string
}

func (c Contract) IsActive(on string) bool {
	return c.EndDate == "" || c.EndDate >= on
}

func (c *Contract) Extend(end string) {
	c.EndDate = end
}
