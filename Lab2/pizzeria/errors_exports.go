package pizzeria

import "lab2/pizzeria/errors"

type (
	EmailNotActiveError      = errors.EmailNotActiveError
	TaskStatusError          = errors.TaskStatusError
	PermissionDeniedError    = errors.PermissionDeniedError
	TeamCapacityError        = errors.TeamCapacityError
	BudgetExceededError      = errors.BudgetExceededError
	CurrencyMismatchError    = errors.CurrencyMismatchError
	TransactionDeclinedError = errors.TransactionDeclinedError
	InventoryShortageError   = errors.InventoryShortageError
	RouteNotFoundError       = errors.RouteNotFoundError
	VehicleOverweightError   = errors.VehicleOverweightError
	OrderValidationError     = errors.OrderValidationError
	SupplierContractError    = errors.SupplierContractError
)
