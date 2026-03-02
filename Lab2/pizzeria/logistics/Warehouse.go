package logistics

import (
	"lab2/pizzeria/errors"
	"lab2/pizzeria/marketing"
)

type Warehouse struct {
	Name     string
	Country  marketing.Country
	Capacity float64
	Used     float64
	Cargo    []Cargo
}

func (w *Warehouse) Store(c Cargo) error {
	if w.Used+c.Weight > w.Capacity {
		return errors.InventoryShortageError{Item: w.Name}
	}
	w.Cargo = append(w.Cargo, c)
	w.Used += c.Weight
	return nil
}

func (w Warehouse) Utilization() float64 {
	if w.Capacity == 0 {
		return 0
	}
	return w.Used / w.Capacity
}
