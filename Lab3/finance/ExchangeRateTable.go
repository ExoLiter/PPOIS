package finance

type ExchangeRateTable struct {
	Base        Currency
	Rates       map[string]float64
	LastUpdated string
}

func NewExchangeRateTable(base Currency) ExchangeRateTable {
	return ExchangeRateTable{Base: base, Rates: map[string]float64{}}
}

func (t *ExchangeRateTable) SetRate(code string, rate float64) {
	if t.Rates == nil {
		t.Rates = map[string]float64{}
	}
	t.Rates[code] = rate
}

func (t ExchangeRateTable) Convert(amount float64, target string) float64 {
	rate, ok := t.Rates[target]
	if !ok || rate == 0 {
		return 0
	}
	return amount * rate
}

func (t ExchangeRateTable) HasRate(code string) bool {
	_, ok := t.Rates[code]
	return ok
}
