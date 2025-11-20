package hr

import (
	"fmt"

	"lab2/pizzeria/errors"
)

type EmailAccount struct {
	Address string
	Active  bool
	Labels  []string
}

func NewEmailAccount(address string) EmailAccount {
	return EmailAccount{Address: address, Labels: []string{"inbox"}}
}

func (e *EmailAccount) Activate() {
	e.Active = true
}

func (e *EmailAccount) AddLabel(label string) {
	e.Labels = append(e.Labels, label)
}

func (e EmailAccount) Send(subject string) (string, error) {
	if !e.Active {
		return "", errors.EmailNotActiveError{Address: e.Address}
	}
	return fmt.Sprintf("email:%s:%s", e.Address, subject), nil
}
