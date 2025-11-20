package performance

import "lab3/hr"

type PromotionCase struct {
	ID           string
	Candidate    hr.Employee
	ProposedRole hr.Position
	Approved     bool
	Decision     string
}

func (c *PromotionCase) Approve() {
	c.Approved = true
	c.Decision = "approved"
	c.Candidate.AssignPosition(c.ProposedRole)
}

func (c *PromotionCase) Deny(reason string) {
	c.Approved = false
	c.Decision = reason
}

func (c PromotionCase) Outcome() string {
	return c.Decision
}
