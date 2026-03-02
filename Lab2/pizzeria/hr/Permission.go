package hr

import "strings"

type Permission struct {
	Code    string
	Scope   string
	Allowed bool
}

func (p *Permission) Allow() {
	p.Allowed = true
}

func (p Permission) Check(code string) bool {
	return p.Allowed && strings.EqualFold(p.Code, code)
}
