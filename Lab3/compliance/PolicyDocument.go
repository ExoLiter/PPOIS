package compliance

type PolicyDocument struct {
	Name    string
	Version string
	Active  bool
}

func (p *PolicyDocument) Publish(version string) {
	p.Version = version
	p.Active = true
}

func (p *PolicyDocument) Retire() {
	p.Active = false
}

func (p PolicyDocument) IsActive() bool {
	return p.Active
}
