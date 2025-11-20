package errors

import "fmt"

type TeamCapacityError struct {
	Team string
}

func (e TeamCapacityError) Error() string {
	return fmt.Sprintf("team %s has no free capacity", e.Team)
}
