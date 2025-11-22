package reporting

type PayrollReport struct {
	Period      string
	Total       float64
	GeneratedBy string
	Items       []float64
}

func (r *PayrollReport) AddItem(amount float64) {
	r.Items = append(r.Items, amount)
}

func (r *PayrollReport) ComputeTotal() {
	total := 0.0
	for _, item := range r.Items {
		total += item
	}
	r.Total = total
}

func (r PayrollReport) HasData() bool {
	return len(r.Items) > 0
}
