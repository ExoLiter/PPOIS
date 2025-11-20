package sales

import (
	"lab2/pizzeria/errors"
	"lab2/pizzeria/logistics"
)

type ProcurementDepartment struct {
	Supplier  Supplier
	Contracts []Contract
	Orders    []PurchaseOrder
	Materials []logistics.Material
}

func (p *ProcurementDepartment) ApproveContract(contract Contract) error {
	if contract.Supplier.Name == "" {
		return errors.SupplierContractError{Supplier: "unknown"}
	}
	p.Contracts = append(p.Contracts, contract)
	return nil
}

func (p ProcurementDepartment) OutstandingOrders() int {
	return len(p.Orders)
}
