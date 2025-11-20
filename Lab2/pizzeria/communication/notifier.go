package communication

import (
	"fmt"
	"sync"
)

type Notifier interface {
	Notify(recipient, message string) error
}

type EmailNotifier struct {
	mu   sync.Mutex
	sent []string
}

func NewEmailNotifier() *EmailNotifier {
	return &EmailNotifier{}
}

func (n *EmailNotifier) Notify(recipient, message string) error {
	if recipient == "" {
		return fmt.Errorf("notifier: empty recipient")
	}
	n.mu.Lock()
	defer n.mu.Unlock()
	n.sent = append(n.sent, fmt.Sprintf("%s:%s", recipient, message))
	return nil
}

func (n *EmailNotifier) Sent() []string {
	n.mu.Lock()
	defer n.mu.Unlock()
	out := make([]string, len(n.sent))
	copy(out, n.sent)
	return out
}
