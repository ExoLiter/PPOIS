package logistics

import "fmt"

type Shipment struct {
	ID      string
	Cargo   []Cargo
	Route   Route
	Vehicle Vehicle
}

func (s *Shipment) AddCargo(c Cargo) {
	s.Cargo = append(s.Cargo, c)
}

func (s Shipment) Summary() string {
	return fmt.Sprintf("shipment:%s:%d", s.ID, len(s.Cargo))
}
