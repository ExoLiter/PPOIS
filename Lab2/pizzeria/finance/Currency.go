package finance

import "fmt"

type Currency struct {
	Code   string
	Symbol string
	Rate   float64
}

func (c Currency) Format(amount float64) string {
	return fmt.Sprintf("%s%.2f", c.Symbol, amount)
}

func (c *Currency) AdjustRate(newRate float64) {
	if newRate > 0 {
		c.Rate = newRate
	}
}
