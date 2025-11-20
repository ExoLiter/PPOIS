package production

import "lab2/pizzeria/logistics"

type TechnologicalCard struct {
	Name    string
	Product logistics.Product
	Steps   []string
}

func (t *TechnologicalCard) AddStep(step string) {
	t.Steps = append(t.Steps, step)
}

func (t TechnologicalCard) StepCount() int {
	return len(t.Steps)
}
