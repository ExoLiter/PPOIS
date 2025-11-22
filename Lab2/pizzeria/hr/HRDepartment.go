package hr

type HRDepartment struct {
	Generator IDGenerator
	Employees []Employee
	Policies  []Permission
}

func (hr *HRDepartment) Hire(name, role string, email EmailAccount) Employee {
	id := hr.Generator.Next()
	emp := Employee{ID: id, Name: name, Role: role, Email: email}
	hr.Employees = append(hr.Employees, emp)
	return emp
}

func (hr *HRDepartment) RevokePermission(code string) int {
	removed := 0
	for i := range hr.Employees {
		filtered := hr.Employees[i].Permissions[:0]
		for _, perm := range hr.Employees[i].Permissions {
			if perm.Check(code) {
				removed++
				continue
			}
			filtered = append(filtered, perm)
		}
		hr.Employees[i].Permissions = filtered
	}
	return removed
}
