package communication

import "testing"

func TestNotificationPreference(t *testing.T) {
	pref := NotificationPreference{EmployeeID: "E1", Channel: "email"}
	pref.Disable()
	if pref.Enabled {
		t.Fatalf("should be disabled")
	}
	pref.Enable()
	pref.UpdateFrequency("daily")
	if !pref.Enabled || pref.Frequency != "daily" {
		t.Fatalf("preference not updated")
	}
}

func TestEmailTemplate(t *testing.T) {
	template := EmailTemplate{Name: "payroll", Subject: "Payroll", Body: "Hi"}
	template.Activate()
	if !template.IsActive() {
		t.Fatalf("template should be active")
	}
	if template.Render("Alice") != "Payroll:Alice" {
		t.Fatalf("render mismatch")
	}
}

func TestMessageSender(t *testing.T) {
	sender := MessageSender{SenderID: "S1", Channel: "email"}
	if !sender.CanSend() {
		t.Fatalf("should be able to send")
	}
	if !sender.Send("hello") {
		t.Fatalf("send failed")
	}
	if sender.SentCount != 1 {
		t.Fatalf("sent count mismatch")
	}
	if sender.Send("") {
		t.Fatalf("empty message should fail")
	}
}

func TestAlertSubscription(t *testing.T) {
	sub := AlertSubscription{ID: "A1", EmployeeID: "E2", Topic: "payroll"}
	sub.Subscribe()
	if !sub.IsActive() {
		t.Fatalf("subscription should be active")
	}
	sub.Unsubscribe()
	if sub.IsActive() {
		t.Fatalf("subscription should be inactive")
	}
}

func TestReminderSchedule(t *testing.T) {
	schedule := ReminderSchedule{ID: "R1"}
	schedule.AddTime("09:00")
	schedule.AddTime("12:00")
	if schedule.NextTime() != "09:00" {
		t.Fatalf("next time mismatch")
	}
	schedule.Toggle()
	if !schedule.Enabled {
		t.Fatalf("toggle failed")
	}
	schedule.RecordSend("09:00")
	if schedule.LastSent != "09:00" {
		t.Fatalf("last sent mismatch")
	}
}
