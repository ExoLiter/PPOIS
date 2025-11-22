package logistics

import (
	"lab2/pizzeria/errors"
	"lab2/pizzeria/marketing"
)

type LogisticsDepartment struct {
	Country   marketing.Country
	Centers   []LogisticsCenter
	Carriers  []Carrier
	Shipments []Shipment
}

func (l *LogisticsDepartment) Dispatch(shipment Shipment) error {
	if shipment.Route.Distance <= 0 {
		return errors.RouteNotFoundError{Route: shipment.Route.Name}
	}
	l.Shipments = append(l.Shipments, shipment)
	return nil
}

func (l LogisticsDepartment) TotalShipments() int {
	return len(l.Shipments)
}
