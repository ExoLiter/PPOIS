package finance

import "fmt"

const defaultPrecision = 2

type Currency struct {
	Code      string
	Symbol    string
	Precision int
}

func (c Currency) Format(amount float64) string {
	return c.Symbol + formatAmount(amount, c.Precision)
}

func (c Currency) Equals(code string) bool {
	return c.Code == code
}

func formatAmount(amount float64, precision int) string {
	scale := 1.0
	for i := 0; i < precision; i++ {
		scale *= 10
	}
	value := float64(int(amount*scale+0.5)) / scale
	return fmtAmount(value, precision)
}

func fmtAmount(amount float64, precision int) string {
	format := "%." + fmt.Sprint(precision) + "f"
	return fmt.Sprintf(format, amount)
}
