package sales

import (
	"lab2/pizzeria/errors"
	"lab2/pizzeria/hr"
)

type Order struct {
	Number   string
	Items    []OrderItem
	Status   hr.Status
	Customer Customer
}

func (o *Order) AddItem(item OrderItem) error {
	if item.Quantity <= 0 {
		return errors.OrderValidationError{Order: o.Number}
	}
	o.Items = append(o.Items, item)
	return nil
}

func (o Order) Total() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.Cost()
	}
	return total
}
