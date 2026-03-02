package sales

import "lab2/pizzeria/hr"

type Customer struct {
	Name   string
	Orders []Order
	Email  hr.EmailAccount
}

func (c *Customer) PlaceOrder(order Order) {
	c.Orders = append(c.Orders, order)
}

func (c Customer) OrderCount() int {
	return len(c.Orders)
}
