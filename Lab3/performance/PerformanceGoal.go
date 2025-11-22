package performance

const goalCompleted = "completed"
const goalInProgress = "in_progress"

type PerformanceGoal struct {
	ID         string
	EmployeeID string
	Target     float64
	Achieved   float64
	Status     string
}

func NewPerformanceGoal(id string, employeeID string, target float64) PerformanceGoal {
	return PerformanceGoal{ID: id, EmployeeID: employeeID, Target: target, Status: goalInProgress}
}

func (g *PerformanceGoal) MarkProgress(value float64) {
	g.Achieved += value
	if g.Achieved >= g.Target {
		g.Status = goalCompleted
	}
}

func (g PerformanceGoal) Completion() float64 {
	if g.Target == 0 {
		return 0
	}
	return g.Achieved / g.Target
}
