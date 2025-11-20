package logistics

import "lab2/pizzeria/errors"

type Vehicle struct {
	Plate string
	Type  VehicleType
	Cargo []Cargo
}

func (v *Vehicle) Load(c Cargo) error {
	weight := 0.0
	for _, existing := range v.Cargo {
		weight += existing.Weight
	}
	if weight+c.Weight > v.Type.Capacity {
		return errors.VehicleOverweightError{Plate: v.Plate}
	}
	v.Cargo = append(v.Cargo, c)
	return nil
}

func (v Vehicle) LoadCount() int {
	return len(v.Cargo)
}
