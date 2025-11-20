package compliance

type RiskAssessment struct {
	ID        string
	Risks     map[string]int
	Threshold int
	Score     int
}

func NewRiskAssessment(id string, threshold int) RiskAssessment {
	return RiskAssessment{ID: id, Risks: map[string]int{}, Threshold: threshold}
}

func (r *RiskAssessment) AddRisk(name string, score int) {
	if r.Risks == nil {
		r.Risks = map[string]int{}
	}
	r.Risks[name] = score
	r.Score += score
}

func (r RiskAssessment) IsCritical() bool {
	return r.Score >= r.Threshold
}
