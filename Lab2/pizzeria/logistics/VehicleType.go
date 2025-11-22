package logistics

type VehicleType struct {
	Name     string
	Capacity float64
	Electric bool
}

func (v VehicleType) CanCarry(weight float64) bool {
	return v.Capacity >= weight
}

func (v *VehicleType) ToggleElectric() {
	v.Electric = !v.Electric
}
