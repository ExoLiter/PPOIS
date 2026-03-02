package finance

import "lab2/pizzeria/errors"

type Budget struct {
	Name       string
	Limit      float64
	Spent      float64
	Categories map[string]float64
}

func NewBudget(name string, limit float64) Budget {
	return Budget{Name: name, Limit: limit, Categories: map[string]float64{}}
}

func (b *Budget) Allocate(category string, amount float64) error {
	if b.Spent+amount > b.Limit {
		return errors.BudgetExceededError{Category: category}
	}
	b.Spent += amount
	b.Categories[category] += amount
	return nil
}

func (b Budget) Remaining() float64 {
	return b.Limit - b.Spent
}
