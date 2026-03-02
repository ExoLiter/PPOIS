package performance

type Certification struct {
	Name       string
	Issuer     string
	ValidUntil string
	Active     bool
}

func (c *Certification) Activate() {
	c.Active = true
}

func (c *Certification) Expire(date string) {
	c.ValidUntil = date
	c.Active = false
}

func (c Certification) IsValid(on string) bool {
	return c.Active && c.ValidUntil >= on
}
