package logistics

import "lab2/pizzeria/marketing"

type LogisticsCenter struct {
	Name    string
	Country marketing.Country
	Routes  []Route
	Plans   []CargoSortingPlan
}

func (l *LogisticsCenter) AddRoute(route Route) {
	l.Routes = append(l.Routes, route)
}

func (l *LogisticsCenter) PlanCargo(plan CargoSortingPlan) {
	l.Plans = append(l.Plans, plan)
}
