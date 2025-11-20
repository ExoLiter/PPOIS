package compliance

type AuditLog struct {
	Actor     string
	Action    string
	Timestamp string
	Entries   []string
}

func (l *AuditLog) AddEntry(entry string) {
	l.Entries = append(l.Entries, entry)
}

func (l AuditLog) EntryCount() int {
	return len(l.Entries)
}

func (l AuditLog) LastEntry() string {
	if len(l.Entries) == 0 {
		return ""
	}
	return l.Entries[len(l.Entries)-1]
}
