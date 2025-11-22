package errors

import "fmt"

type SupplierContractError struct {
	Supplier string
}

func (e SupplierContractError) Error() string {
	return fmt.Sprintf("supplier %s has no valid contract", e.Supplier)
}
