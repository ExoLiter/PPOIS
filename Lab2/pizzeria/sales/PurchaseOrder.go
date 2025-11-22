package sales

import "lab2/pizzeria/logistics"

type PurchaseOrder struct {
	Number    string
	Supplier  Supplier
	Materials []logistics.Material
	Products  []logistics.Product
}

func (p *PurchaseOrder) AddMaterial(material logistics.Material) {
	p.Materials = append(p.Materials, material)
}

func (p PurchaseOrder) MaterialCount() int {
	return len(p.Materials)
}
