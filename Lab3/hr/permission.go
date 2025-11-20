package hr

type Permission struct {
	Name    string
	Scope   string
	Allowed bool
}

func (p *Permission) Grant() {
	p.Allowed = true
}

func (p *Permission) Revoke() {
	p.Allowed = false
}

func (p Permission) CanAccess(scope string) bool {
	return p.Allowed && p.Scope == scope
}
