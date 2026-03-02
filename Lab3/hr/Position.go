package hr

const levelStep = 1
const baseRaise = 50.0

type Position struct {
	Title      string
	Level      int
	BaseSalary float64
}

func (p *Position) Promote() {
	p.Level += levelStep
	p.BaseSalary += baseRaise
}

func (p Position) IsSenior() bool {
	return p.Level >= 3
}
