package benefits

type BenefitPlan struct {
	Name                string
	Provider            string
	Cost                float64
	EmployerContribution float64
	Active              bool
}

func (p *BenefitPlan) Activate() {
	p.Active = true
}

func (p BenefitPlan) EmployerShare() float64 {
	return p.Cost * p.EmployerContribution
}

func (p BenefitPlan) EmployeeShare() float64 {
	share := p.Cost - p.EmployerShare()
	if share < 0 {
		return 0
	}
	return share
}
