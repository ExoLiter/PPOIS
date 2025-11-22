package compliance

const severityHigh = "high"

type IncidentReport struct {
	ID          string
	Severity    string
	Description string
	Resolved    bool
	Resolution  string
}

func (r *IncidentReport) Resolve(resolution string) {
	r.Resolution = resolution
	r.Resolved = true
}

func (r IncidentReport) IsHighSeverity() bool {
	return r.Severity == severityHigh
}
