package finance

import "lab2/pizzeria/errors"

type CurrencyConverter struct {
	Base  Currency
	Rates map[string]float64
}

func NewCurrencyConverter(base Currency) CurrencyConverter {
	return CurrencyConverter{Base: base, Rates: map[string]float64{base.Code: 1}}
}

func (c *CurrencyConverter) AddRate(code string, rate float64) {
	if c.Rates == nil {
		c.Rates = map[string]float64{}
	}
	c.Rates[code] = rate
}

func (c CurrencyConverter) Convert(amount float64, to Currency) (float64, error) {
	rate, ok := c.Rates[to.Code]
	if !ok {
		return 0, errors.CurrencyMismatchError{From: c.Base.Code, To: to.Code}
	}
	return amount * rate, nil
}
