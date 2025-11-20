package errors

import "fmt"

type RouteNotFoundError struct {
	Route string
}

func (e RouteNotFoundError) Error() string {
	return fmt.Sprintf("route %s is invalid", e.Route)
}
