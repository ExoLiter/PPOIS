package errors

import "fmt"

type VehicleOverweightError struct {
	Plate string
}

func (e VehicleOverweightError) Error() string {
	return fmt.Sprintf("vehicle %s overloaded", e.Plate)
}
