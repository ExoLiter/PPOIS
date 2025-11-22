package logistics

// CargoSortingPlan ties cargo batches to a specific warehouse schedule.
type CargoSortingPlan struct {
	Name      string
	Warehouse *Warehouse
	Items     []Cargo
}

func (c *CargoSortingPlan) AddCargo(cargo Cargo) {
	c.Items = append(c.Items, cargo)
}

func (c CargoSortingPlan) TotalWeight() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.Weight
	}
	return total
}

// WarehouseCapacity returns the warehouse capacity left for the plan, if warehouse attached.
func (c CargoSortingPlan) WarehouseCapacity() float64 {
	if c.Warehouse == nil {
		return 0
	}
	return c.Warehouse.Capacity - c.Warehouse.Used
}
