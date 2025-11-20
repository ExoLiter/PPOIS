package production

import "lab2/pizzeria/logistics"

type ProductionUnit struct {
	Name     string
	Product  logistics.Product
	Capacity int
	Active   bool
	Orders   []ProductionOrder
}

func (u *ProductionUnit) AssignOrder(order ProductionOrder) {
	u.Orders = append(u.Orders, order)
}

func (u *ProductionUnit) Toggle(active bool) {
	u.Active = active
}
