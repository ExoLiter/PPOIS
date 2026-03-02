package production

import "lab2/pizzeria/logistics"

type Factory struct {
	Name       string
	Warehouses []*logistics.Warehouse
	Lines      []ProductionLine
	Plan       ProductionPlan
	Orders     []ProductionOrder
}

func (f *Factory) AddLine(line ProductionLine) {
	f.Lines = append(f.Lines, line)
}

func (f *Factory) AttachWarehouse(wh *logistics.Warehouse) {
	f.Warehouses = append(f.Warehouses, wh)
}

func (f *Factory) Produce() int {
	total := 0
	for _, line := range f.Lines {
		total += line.Load()
	}
	return total
}
