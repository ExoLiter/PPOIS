package compliance

type AccessReview struct {
	ID       string
	Reviewer string
	Findings []string
	Approved bool
}

func (r *AccessReview) AddFinding(finding string) {
	r.Findings = append(r.Findings, finding)
}

func (r *AccessReview) Approve() {
	r.Approved = true
}

func (r AccessReview) HasFindings() bool {
	return len(r.Findings) > 0
}
