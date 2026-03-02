package logistics

import "lab2/pizzeria/errors"

type Material struct {
	Name  string
	Unit  string
	Stock float64
}

func (m *Material) Consume(amount float64) error {
	if amount > m.Stock {
		return errors.InventoryShortageError{Item: m.Name}
	}
	m.Stock -= amount
	return nil
}

func (m *Material) Restock(amount float64) {
	m.Stock += amount
}
