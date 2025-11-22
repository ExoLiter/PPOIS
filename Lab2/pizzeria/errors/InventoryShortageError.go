package errors

import "fmt"

type InventoryShortageError struct {
	Item string
}

func (e InventoryShortageError) Error() string {
	return fmt.Sprintf("inventory shortage for %s", e.Item)
}
