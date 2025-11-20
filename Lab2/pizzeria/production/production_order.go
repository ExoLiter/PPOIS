package production

import "lab2/pizzeria/logistics"

type ProductionOrder struct {
	Number   string
	Card     TechnologicalCard
	Product  logistics.Product
	Quantity int
	Planned  int
}

func (p *ProductionOrder) Schedule(amount int) {
	p.Planned += amount
}

func (p ProductionOrder) Remaining() int {
	if p.Planned < p.Quantity {
		return p.Quantity - p.Planned
	}
	return 0
}
