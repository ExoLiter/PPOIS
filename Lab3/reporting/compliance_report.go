package reporting

type ComplianceReport struct {
	Issues    []string
	RiskLevel int
	Reviewer  string
}

func (r *ComplianceReport) AddIssue(issue string, risk int) {
	r.Issues = append(r.Issues, issue)
	r.RiskLevel += risk
}

func (r ComplianceReport) IsClean() bool {
	return len(r.Issues) == 0
}
