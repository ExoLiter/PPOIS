package hr

const bonusIncrease = 100

type Employee struct {
	ID         string
	Name       string
	Position   Position
	Contract   Contract
	Salary     float64
	Department string
	Active     bool
}

func (e *Employee) Activate() {
	e.Active = true
}

func (e *Employee) AssignPosition(position Position) {
	e.Position = position
}

func (e *Employee) UpdateSalary(delta float64) {
	e.Salary += delta
	if e.Salary < 0 {
		e.Salary = 0
	}
}

func (e *Employee) AwardBonus() {
	e.Salary += bonusIncrease
}
