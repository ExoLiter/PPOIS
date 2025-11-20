package hr

type AccessBadge struct {
	ID         string
	EmployeeID string
	Active     bool
	LastUsed   string
	Location   string
}

func (b *AccessBadge) ActivateBadge() {
	b.Active = true
}

func (b *AccessBadge) UseBadge(location string) string {
	if !b.Active {
		return ""
	}
	b.LastUsed = location
	b.Location = location
	return location
}

func (b AccessBadge) MatchesUser(user string) bool {
	return b.EmployeeID == user
}
