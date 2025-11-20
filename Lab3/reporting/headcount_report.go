package reporting

type HeadcountReport struct {
	DepartmentCounts map[string]int
	ActiveEmployees  int
}

func NewHeadcountReport() HeadcountReport {
	return HeadcountReport{DepartmentCounts: map[string]int{}}
}

func (r *HeadcountReport) AddDepartment(name string, count int) {
	if r.DepartmentCounts == nil {
		r.DepartmentCounts = map[string]int{}
	}
	r.DepartmentCounts[name] = count
	r.ActiveEmployees += count
}

func (r HeadcountReport) Total() int {
	return r.ActiveEmployees
}

func (r HeadcountReport) DepartmentCount(name string) int {
	return r.DepartmentCounts[name]
}
