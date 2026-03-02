package hr

import "lab2/pizzeria/errors"

type Team struct {
	Name     string
	Members  []Employee
	Capacity int
}

func (t *Team) AddMember(emp Employee) error {
	if len(t.Members) >= t.Capacity {
		return errors.TeamCapacityError{Team: t.Name}
	}
	t.Members = append(t.Members, emp)
	return nil
}

func (t Team) Utilization() float64 {
	if t.Capacity == 0 {
		return 0
	}
	return float64(len(t.Members)) / float64(t.Capacity)
}
