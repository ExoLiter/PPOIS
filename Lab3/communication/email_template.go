package communication

type EmailTemplate struct {
	Name    string
	Subject string
	Body    string
	Active  bool
}

func (t *EmailTemplate) Activate() {
	t.Active = true
}

func (t EmailTemplate) Render(name string) string {
	return t.Subject + ":" + name
}

func (t EmailTemplate) IsActive() bool {
	return t.Active
}
