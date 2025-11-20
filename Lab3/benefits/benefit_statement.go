package benefits

type BenefitStatement struct {
	EmployeeID string
	Plans      []BenefitPlan
	TotalCost  float64
}

func (s *BenefitStatement) AddPlan(plan BenefitPlan) {
	s.Plans = append(s.Plans, plan)
}

func (s *BenefitStatement) ComputeTotal() {
	total := 0.0
	for _, plan := range s.Plans {
		total += plan.Cost
	}
	s.TotalCost = total
}

func (s BenefitStatement) PlanCount() int {
	return len(s.Plans)
}
