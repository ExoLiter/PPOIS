package sales

import "lab2/pizzeria/logistics"

type Delivery struct {
	Number    string
	Order     Order
	Route     logistics.Route
	Delivered bool
}

func (d *Delivery) Dispatch(route logistics.Route) {
	d.Route = route
}

func (d *Delivery) Complete() {
	d.Delivered = true
}
