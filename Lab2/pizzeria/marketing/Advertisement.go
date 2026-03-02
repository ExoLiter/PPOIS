package marketing

import "fmt"

type Advertisement struct {
	Title    string
	Message  string
	Channels []string
}

func (a *Advertisement) AddChannel(name string) {
	a.Channels = append(a.Channels, name)
}

func (a Advertisement) Content() string {
	return fmt.Sprintf("%s:%s", a.Title, a.Message)
}
