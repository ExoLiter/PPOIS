package hr

import "strings"

type Status struct {
	Name    string
	Level   int
	Details []string
}

func (s *Status) Update(name string, level int) {
	s.Name = name
	s.Level = level
}

func (s *Status) Annotate(note string) {
	s.Details = append(s.Details, note)
}

func (s Status) IsCritical() bool {
	return strings.EqualFold(s.Name, "critical") || s.Level >= 5
}
