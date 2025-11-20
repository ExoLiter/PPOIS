package hr

import "fmt"

type IDGenerator struct {
	Prefix  string
	counter int
	issued  map[string]struct{}
}

func NewIDGenerator(prefix string) IDGenerator {
	return IDGenerator{Prefix: prefix, issued: map[string]struct{}{}, counter: 0}
}

func (g *IDGenerator) Next() string {
	g.counter++
	id := fmt.Sprintf("%s-%03d", g.Prefix, g.counter)
	g.issued[id] = struct{}{}
	return id
}

func (g *IDGenerator) Reset() {
	g.issued = map[string]struct{}{}
}
