package errors_test

import (
	"testing"

	"lab2/pizzeria/errors"
)

func TestErrorMessagesNotEmpty(t *testing.T) {
	errs := []error{
		errors.EmailNotActiveError{Address: "a@pizza.io"},
		errors.TaskStatusError{Task: "task"},
		errors.PermissionDeniedError{Code: "perm"},
		errors.TeamCapacityError{Team: "Kitchen"},
		errors.BudgetExceededError{Category: "ops"},
		errors.CurrencyMismatchError{From: "USD", To: "EUR"},
		errors.TransactionDeclinedError{Reason: "declined"},
		errors.InventoryShortageError{Item: "cheese"},
		errors.RouteNotFoundError{Route: "R1"},
		errors.VehicleOverweightError{Plate: "V1"},
		errors.OrderValidationError{Order: "O1"},
		errors.SupplierContractError{Supplier: "S1"},
	}

	for i, err := range errs {
		if msg := err.Error(); msg == "" {
			t.Fatalf("error %d produced empty message", i)
		}
	}
}
