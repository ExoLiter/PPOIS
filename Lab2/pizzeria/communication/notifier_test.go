package communication_test

import (
	"testing"

	"lab2/pizzeria/communication"
)

func TestEmailNotifier(t *testing.T) {
	notifier := communication.NewEmailNotifier()

	if err := notifier.Notify("", "welcome"); err == nil {
		t.Fatalf("expected validation error for empty recipient")
	}
	if err := notifier.Notify("chef@pizza.io", "shift"); err != nil {
		t.Fatalf("notify returned error: %v", err)
	}
	if sent := notifier.Sent(); len(sent) != 1 {
		t.Fatalf("expected one notification, got %d", len(sent))
	}
}
