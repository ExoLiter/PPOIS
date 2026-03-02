package logistics

import "lab2/pizzeria/errors"

type Route struct {
	Name     string
	Stops    []string
	Distance float64
}

func (r *Route) AddStop(stop string) {
	r.Stops = append(r.Stops, stop)
}

func (r Route) TravelTime(speed float64) (float64, error) {
	if speed <= 0 {
		return 0, errors.RouteNotFoundError{Route: r.Name}
	}
	return r.Distance / speed, nil
}
