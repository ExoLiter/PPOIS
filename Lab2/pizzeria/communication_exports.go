package pizzeria

import "lab2/pizzeria/communication"

type (
	Notifier      = communication.Notifier
	EmailNotifier = communication.EmailNotifier
)

var NewEmailNotifier = communication.NewEmailNotifier
