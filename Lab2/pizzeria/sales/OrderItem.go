package sales

import (
	"fmt"

	"lab2/pizzeria/logistics"
)

type OrderItem struct {
	Product  logistics.Product
	Quantity int
	Price    float64
}

func (i OrderItem) Cost() float64 {
	return float64(i.Quantity) * i.Price
}

func (i OrderItem) Describe() string {
	return fmt.Sprintf("%s x%d", i.Product.Name, i.Quantity)
}
