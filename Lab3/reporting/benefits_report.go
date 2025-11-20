package reporting

type BenefitsReport struct {
	EmployeeID   string
	BenefitTotal float64
	PlanCount    int
}

func (r *BenefitsReport) AddBenefit(cost float64) {
	r.BenefitTotal += cost
	r.PlanCount++
}

func (r BenefitsReport) AverageCost() float64 {
	if r.PlanCount == 0 {
		return 0
	}
	return r.BenefitTotal / float64(r.PlanCount)
}
