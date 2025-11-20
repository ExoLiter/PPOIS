package logistics

type Carrier struct {
	Name      string
	Vehicles  []Vehicle
	Shipments []Shipment
}

func (c *Carrier) AssignVehicle(vehicle Vehicle) {
	c.Vehicles = append(c.Vehicles, vehicle)
}

func (c *Carrier) Schedule(shipment Shipment) {
	c.Shipments = append(c.Shipments, shipment)
}
