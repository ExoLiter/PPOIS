package production

import "lab2/pizzeria/logistics"

type ProductionPlan struct {
	Name      string
	Orders    []ProductionOrder
	Materials []logistics.Material
}

func (p *ProductionPlan) AddOrder(order ProductionOrder) {
	p.Orders = append(p.Orders, order)
}

func (p ProductionPlan) NeedsMaterial(name string) bool {
	for _, material := range p.Materials {
		if material.Name == name {
			return true
		}
	}
	return false
}
