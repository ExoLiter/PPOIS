package communication

type MessageSender struct {
	SenderID  string
	SentCount int
	Channel   string
}

func (s *MessageSender) Send(message string) bool {
	if message == "" {
		return false
	}
	s.SentCount++
	return true
}

func (s MessageSender) CanSend() bool {
	return s.Channel != ""
}
