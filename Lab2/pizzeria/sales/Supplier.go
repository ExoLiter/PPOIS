package sales

import (
	"lab2/pizzeria/errors"
	"lab2/pizzeria/logistics"
	"lab2/pizzeria/marketing"
)

type Supplier struct {
	Name      string
	Country   marketing.Country
	Materials []logistics.Material
	Products  []logistics.Product
}

func (s *Supplier) ProvideMaterial(name string) (*logistics.Material, error) {
	for i, material := range s.Materials {
		if material.Name == name {
			return &s.Materials[i], nil
		}
	}
	return nil, errors.SupplierContractError{Supplier: s.Name}
}

func (s *Supplier) ProvideProduct(name string) (*logistics.Product, error) {
	for i, product := range s.Products {
		if product.Name == name {
			return &s.Products[i], nil
		}
	}
	return nil, errors.SupplierContractError{Supplier: s.Name}
}
