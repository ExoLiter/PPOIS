package payroll

import "lab3/hr"

type PayrollRun struct {
	ID        string
	Period    string
	Employees []hr.Employee
	Calendar  PayCalendar
	Components []PayComponent
}

func (r *PayrollRun) AddEmployee(e hr.Employee) {
	r.Employees = append(r.Employees, e)
}

func (r *PayrollRun) AddComponent(c PayComponent) {
	r.Components = append(r.Components, c)
}

func (r PayrollRun) TotalPayroll() float64 {
	total := 0.0
	for _, e := range r.Employees {
		amount := e.Salary
		for _, comp := range r.Components {
			amount = comp.ApplyTo(amount)
		}
		total += amount
	}
	return total
}

func (r PayrollRun) UsesCalendar() bool {
	return r.Calendar.IsValid()
}
