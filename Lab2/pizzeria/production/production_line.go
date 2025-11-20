package production

import "lab2/pizzeria/logistics"

type ProductionLine struct {
	Name    string
	Unit    ProductionUnit
	Orders  []ProductionOrder
	Outputs []logistics.Product
}

func (p *ProductionLine) Enqueue(order ProductionOrder) {
	p.Orders = append(p.Orders, order)
}

func (p *ProductionLine) Load() int {
	total := 0
	for _, order := range p.Orders {
		total += order.Quantity
	}
	return total
}
